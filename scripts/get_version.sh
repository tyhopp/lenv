#!/bin/sh

git fetch --tags
VERSION=$(git describe --tags `git rev-list --tags --max-count=1` || echo "none")

if [ "$VERSION" = "none" ]; then
  echo "No tags found, exiting"
  exit 0
fi

echo "VERSION=$VERSION" > .version