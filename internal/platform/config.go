package platform

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Address                   string
	ReadTimeout, WriteTimeout time.Duration
	RequestLimit              int64
}

func LoadConfig(path string) Config {
	c := Config{Address: ":8088", ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, RequestLimit: 1 << 20}
	if f, err := os.Open(path); err == nil {
		s := bufio.NewScanner(f)
		for s.Scan() {
			p := strings.SplitN(strings.TrimSpace(s.Text()), ":", 2)
			if len(p) != 2 {
				continue
			}
			k := strings.TrimSpace(p[0])
			v := strings.Trim(strings.TrimSpace(p[1]), "\"'")
			switch k {
			case "address":
				c.Address = v
			case "read_timeout":
				if d, e := time.ParseDuration(v); e == nil {
					c.ReadTimeout = d
				}
			case "write_timeout":
				if d, e := time.ParseDuration(v); e == nil {
					c.WriteTimeout = d
				}
			case "request_limit":
				if n, e := strconv.ParseInt(v, 10, 64); e == nil {
					c.RequestLimit = n
				}
			}
		}
		_ = f.Close()
	}
	if v := os.Getenv("REGISTRY_SERVER_ADDRESS"); v != "" {
		c.Address = v
	}
	return c
}
