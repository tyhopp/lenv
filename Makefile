.PHONY: all setup deps build build-all

all: build-all

setup:
	go mod tidy

deps:
	go get .

build: setup deps
	sh ./scripts/build.sh

build-all: setup deps
	sh ./scripts/build_all.sh
