#!/bin/sh

for goos in linux windows; do
  for goarch in amd64 arm64; do
    GOOS=$goos GOARCH=$goarch sh ./scripts/build.sh
  done
done