#!/bin/sh

if [ -z "$GOOS" ] || [ -z "$GOARCH" ]; then
  echo "GOOS and/or GOARCH not set, exiting"
  exit 1
fi

EXT=""
if [ "$GOOS" = "windows" ]; then
  EXT=".exe"
fi

go build -ldflags="-s -w" -trimpath -o lenv-$GOOS-$GOARCH$EXT ./cmd/lenv/main.go