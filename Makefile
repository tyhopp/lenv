.PHONY: all deps build

all: build

deps:
	go mod tidy
	go get .

build: deps
	go build -ldflags="-s -w" -trimpath -o lenv-linux-amd64 ./cmd/lenv/main.go
	go build -ldflags="-s -w" -trimpath -o lenv-linux-arm64 ./cmd/lenv/main.go
	go build -ldflags="-s -w" -trimpath -o lenv-windows-amd64.exe ./cmd/lenv/main.go
	go build -ldflags="-s -w" -trimpath -o lenv-windows-arm64.exe ./cmd/lenv/main.go
