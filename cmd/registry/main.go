package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/example/api-schema-registry/internal/asset"
	"github.com/example/api-schema-registry/internal/compatibility"
	"github.com/example/api-schema-registry/internal/consumer"
	"github.com/example/api-schema-registry/internal/distribution"
	"github.com/example/api-schema-registry/internal/lineage"
	"github.com/example/api-schema-registry/internal/namespace"
	"github.com/example/api-schema-registry/internal/platform"
	"github.com/example/api-schema-registry/internal/validation"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type app struct {
	ns        *namespace.Service
	assets    *asset.Service
	arepo     *asset.Memory
	consumers *consumer.Service
	crepo     *consumer.Memory
	compat    *compatibility.Service
	validator *validation.Service
	events    *distribution.Service
	graph     *lineage.Service
	log       *slog.Logger
	metrics   *platform.Counters
	ready     *platform.Readiness
}

func main() {
	cfg := platform.LoadConfig(*flag.String("config", "configs/config.yaml", "config path"))
	flag.Parse()
	log := platform.Logger()
	ar := asset.NewMemory()
	cr := consumer.NewMemory()
	g := lineage.New()
	a := &app{ns: namespace.NewService(namespace.NewMemory()), assets: asset.NewService(ar), arepo: ar, consumers: consumer.NewService(cr), crepo: cr, compat: compatibility.NewService(), validator: validation.NewService(), events: distribution.NewService(distribution.New()), graph: lineage.NewService(g), log: log, metrics: &platform.Counters{}, ready: &platform.Readiness{}}
	a.ready.Set(true)
	srv := &http.Server{Addr: cfg.Address, Handler: a.routes(), ReadTimeout: cfg.ReadTimeout, WriteTimeout: cfg.WriteTimeout}
	go func() {
		log.Info("registry started", "address", cfg.Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server stopped", "error", err)
			os.Exit(1)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Info("registry shutdown")
}
func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/healthz", platform.Health(a.ready, a.metrics))
	mux.Handle("/readyz", platform.Health(a.ready, a.metrics))
	mux.Handle("/metrics", platform.Health(a.ready, a.metrics))
	mux.Handle("/v1/namespaces", platform.JSON(a.createNamespace))
	mux.Handle("/v1/assets", platform.JSON(a.assetRoutes))
	mux.Handle("/v1/assets/", platform.JSON(a.assetRoutes))
	mux.Handle("/v1/compatibility", platform.JSON(a.compatibility))
	mux.Handle("/v1/validate", platform.JSON(a.validate))
	mux.Handle("/v1/consumers", platform.JSON(a.createConsumer))
	mux.Handle("/v1/impact", platform.JSON(a.impact))
	mux.Handle("/v1/changes", platform.JSON(a.changes))
	return platform.Recover(platform.RequestID(mux))
}
func (a *app) createNamespace(ctx context.Context, r *http.Request) (any, error) {
	if r.Method != "POST" {
		return nil, fmt.Errorf("method not allowed")
	}
	var in struct {
		Name  string   `json:"name"`
		Owner string   `json:"owner"`
		Tags  []string `json:"tags"`
	}
	if err := platform.Decode(r, &in); err != nil {
		return nil, err
	}
	a.metrics.IncRequest()
	return a.ns.Create(ctx, in.Name, in.Owner, in.Tags)
}
func (a *app) assetRoutes(ctx context.Context, r *http.Request) (any, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) == 2 && r.Method == "POST" {
		var in struct {
			NamespaceID string   `json:"namespace_id"`
			Name        string   `json:"name"`
			Owner       string   `json:"owner"`
			Kind        string   `json:"kind"`
			Tags        []string `json:"tags"`
		}
		if err := platform.Decode(r, &in); err != nil {
			return nil, err
		}
		return a.assets.Create(ctx, in.NamespaceID, in.Name, in.Owner, asset.Kind(in.Kind), in.Tags)
	}
	if len(parts) >= 6 && parts[2] != "" && parts[3] == "versions" && parts[5] == "publish" && r.Method == "POST" {
		v, e := a.assets.GetVersion(ctx, parts[2], parts[4])
		if e != nil {
			return nil, e
		}
		p := map[string]any{}
		for _, old := range a.arepo.Versions(parts[2]) {
			if old.Status == asset.Published {
				p = old.Document
			}
		}
		if len(p) > 0 {
			res := compatibility.Compare(p, v.Document, compatibility.Backward)
			if !res.Compatible {
				return nil, fmt.Errorf("incompatible: %v", res.Differences)
			}
		}
		if e := a.arepo.SetPublished(parts[2], parts[4]); e != nil {
			return nil, e
		}
		return a.events.Publish(ctx, parts[2], parts[4], map[string]any{"checksum": v.Checksum}), nil
	}
	if len(parts) == 4 && parts[2] != "" && parts[3] == "versions" && r.Method == "POST" {
		var in struct {
			Version  string         `json:"version"`
			Document map[string]any `json:"document"`
		}
		if err := platform.Decode(r, &in); err != nil {
			return nil, err
		}
		return a.assets.CreateVersion(ctx, parts[2], in.Version, in.Document)
	}
	return nil, fmt.Errorf("unknown asset route")
}
func (a *app) compatibility(ctx context.Context, r *http.Request) (any, error) {
	var in struct {
		Old  map[string]any `json:"old"`
		New  map[string]any `json:"new"`
		Mode string         `json:"mode"`
	}
	if err := platform.Decode(r, &in); err != nil {
		return nil, err
	}
	return a.compat.Check(ctx, in.Old, in.New, compatibility.Mode(in.Mode))
}
func (a *app) validate(ctx context.Context, r *http.Request) (any, error) {
	var in struct {
		Schema any `json:"schema"`
		Value  any `json:"value"`
	}
	if err := platform.Decode(r, &in); err != nil {
		return nil, err
	}
	return a.validator.Validate(ctx, in.Schema, in.Value)
}
func (a *app) createConsumer(ctx context.Context, r *http.Request) (any, error) {
	var in struct {
		Name        string `json:"name"`
		AssetID     string `json:"asset_id"`
		Version     string `json:"version"`
		Environment string `json:"environment"`
		Strategy    string `json:"strategy"`
	}
	if err := platform.Decode(r, &in); err != nil {
		return nil, err
	}
	return a.consumers.Register(ctx, in.Name, in.AssetID, in.Version, in.Environment, in.Strategy)
}
func (a *app) impact(ctx context.Context, r *http.Request) (any, error) {
	id := r.URL.Query().Get("asset_id")
	v := r.URL.Query().Get("version")
	cs := a.crepo.ListByAsset(id, v)
	down, _ := a.graph.Impact(ctx, id)
	return map[string]any{"consumers": cs, "downstream": down}, nil
}
func (a *app) changes(ctx context.Context, r *http.Request) (any, error) {
	cur, _ := strconv.ParseInt(r.URL.Query().Get("cursor"), 10, 64)
	ev := a.events.Changes(ctx, cur, 100)
	return map[string]any{"events": ev, "next_cursor": cursor(ev, cur)}, nil
}
func cursor(ev []distribution.Event, cur int64) int64 {
	if len(ev) == 0 {
		return cur
	}
	return ev[len(ev)-1].Sequence
}
func readAll(r *http.Request) ([]byte, error) {
	b, e := io.ReadAll(io.LimitReader(r.Body, 2<<20))
	if e != nil {
		return nil, fmt.Errorf("read body: %w", e)
	}
	return b, nil
}
func decode(data []byte, v any) error { return json.Unmarshal(data, v) }
