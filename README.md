# lenv

Manage symlinks from a root file to multiple destinations.

Useful for monorepos that use a single `.env` file as a source of truth for many child projects.

## Installation

### Via Go tooling

- `go install github.com/tyhopp/lenv/cmd/lenv@latest` to install the command line executable
- `go get github.com/tyhopp/lenv` to install as a dependency in your Go project

### Via precompiled binaries

To install the latest linux amd64 release from GitHub:

```
curl -L https://github.com/tyhopp/lenv/releases/latest/download/lenv-linux-amd64
```

See [Releases](https://github.com/tyhopp/lenv/releases) for all available binaries and versions.

## Usage

In the root of your project:

1. Create a `.env` (or other named) file you want to symlink
2. Create a `.lenv` file with the destination locations to symlink to, such as:

```
project/a/.env
project/b/.env
```

3. Follow these usage instructions:
```
Usage: lenv [options] <subcommand>
Options:
  -env string
        name of the environment file (default ".env")
  -help
        display help information
Subcommands:
  check   - Check status of symlinks between source env file and destinations
  link    - Symlink source env file to destinations
  unlink  - Remove symlinks from destinations
```

### WASI binary execution

The [WASI](https://wasi.dev/) binary can be executed with the [Wasmtime](https://wasmtime.dev/) WebAssembly runtime with this command structure:

```
wasmtime --wasi cli --dir /absolute/path/to/project lenv-wasip1.wasm
```