#!/usr/bin/env sh
set -eu
exec go run ./cmd/registry -config "${REGISTRY_CONFIG:-configs/config.yaml}"
