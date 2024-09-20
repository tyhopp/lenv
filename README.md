# lenv

Manage symlinks from a root file to multiple destinations.

Useful for monorepos that use a single `.env` file as a source of truth for many child projects.

## Installation

### Via Go tooling

- `go install github.com/tyhopp/lenv` to install the command line executable
- `go get github.com/tyhopp/lenv` to install as a dependency in your Go project

### Via precompiled binaries

To install and make executable the latest release from GitHub:

```sh
curl -L \
  -o /usr/local/bin/lenv \
  https://github.com/tyhopp/lenv/releases/latest/download/lenv-linux
chmod +x /usr/local/bin/lenv
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