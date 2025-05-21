# Notes CLI

A simple Go-based CLI tool to manage markdown notes form your terminal.

- Automatically creates and stores notes in `~/notes-app`
- Opens notes in Neovim.
- Lists all `.md` notes when no arguments are passed.

## Features

- Create or open a note with a single argument
- Auto-add `.md` extension if none provided
- List all existing notes
- Opens files in Neovim
- Cross-platform -- works on Linux, macOS, Windows (with Neovim installed)

## Installation

1. Clone the repository

2. Build the binary:
```bash
go build -o notes
```

3. (Optional) Move it to your `$PATH` for global usage:
```bash
sudo mv notes /usr/local/bin/
```

## Usage

| Commands          | Description                           |
| ----------------- | ------------------------------------- |
| notes             | List all `.md` notes                  |
| notes mynote      | Creates/opens `mynote.md`             |
| notes log.txt     | Creates/opens `log.txt`               |
| notes a b         | Shows usage help (too many arguments) |

## Examples

```bash
notes               # List all notes
notes journal       # Create/open journal.md
notes ideas.txt     # Create/open ideas.txt
```

## Notes Directory

All notes are stored in:

```bash
~/notes-app/
```

This folder is created automatically if it doesn't exist.

## Editor

This tool uses **Neovim** (`nvim`) by default.
Ensure `nvim` is installed and accessible in your `$PATH`

## Requirements

- [Go](https://golang.org/doc/install) 1.18+
- [Neovim](https://neovim.io/) installed and in your `$PATH`
