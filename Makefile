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
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o lenv-darwin-amd64 ./cmd/lenv/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o lenv-darwin-arm64 ./cmd/lenv/main.go
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o lenv-freebsd-amd64 ./cmd/lenv/main.go
	GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o lenv-freebsd-arm64 ./cmd/lenv/main.go
	GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o lenv-openbsd-amd64 ./cmd/lenv/main.go
	GOOS=openbsd GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o lenv-openbsd-arm64 ./cmd/lenv/main.go
	GOOS=wasip1 GOARCH=wasm go build -ldflags="-s -w" -trimpath -o lenv-wasip1.wasm ./cmd/lenv/main.go
