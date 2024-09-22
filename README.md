# lenv

Manage symlinks from a root file to multiple destinations.

Useful for monorepos that use a single `.env` file as a source of truth for many child projects.

## Installation

Download precompiled binaries from [Releases](https://github.com/tyhopp/lenv/releases).

Ports for various programming languages are available in [lenv-ports](https://github.com/tyhopp/lenv-ports):

- [JavaScript](https://www.npmjs.com/package/lenv-js) (Node.js, ESM)

## Usage

In the root of your project:

1. Create a `.env` (or other named) file you want to symlink
2. Create a `.lenv` file with the destination locations to symlink to, such as:

```
project/a/.env
project/b/.env
```

3. Execute `lenv link` to create symlinks

Use the `-help` flag to see all usage instructions.

### WASI binary execution

The [WebAssembly System Interface (WASI)](https://wasi.dev/) binary can be executed with the [Wasmtime](https://wasmtime.dev/) runtime with this command structure:

```
wasmtime --wasi cli --dir /absolute/path/to/project lenv-wasip1.wasm
```