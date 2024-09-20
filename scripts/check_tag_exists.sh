#!/bin/sh

VERSION=$(cat .version | cut -d'=' -f2)

if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "Tag $VERSION already exists, exiting"
  exit 0
fi