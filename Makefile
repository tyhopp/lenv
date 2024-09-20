.PHONY: all setup deps build build-all get-version check-tag-exists

all: build-all

setup:
	go mod tidy

deps:
	go get .

build: setup deps
	sh ./scripts/build.sh

build-all: setup deps
	sh ./scripts/build_all.sh

get-version:
	sh ./scripts/get-version.sh

check-tag-exists: get-version
	sh ./scripts/check-tag-exists.sh