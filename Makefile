.PHONY: all deps build

all: build

deps:
	go mod tidy
	go get .

build: deps
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o lenv-linux-amd64 ./cmd/lenv/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o lenv-linux-arm64 ./cmd/lenv/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o lenv-windows-amd64.exe ./cmd/lenv/main.go
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o lenv-windows-arm64.exe ./cmd/lenv/main.go
