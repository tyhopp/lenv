# lenv

Manage symlinks from a root file to multiple destinations.

Useful for monorepos that use a single `.env` file as a source of truth for many child projects.

## Installation

TBD

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