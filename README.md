# Dirtools

Dirtools is a collection of powerful command-line utilities for analyzing and searching directory structures. It provides tools for isolating directories based on file extensions and searching for directories containing files matching specific patterns.

## Installation

```bash
go install github.com/arran4/dirtools/cmd/...@latest
```

## Tools

### extdirisolate

`extdirisolate` is a utility for analyzing directory structures to identify folders containing specific types of files. It is particularly useful for organizing media libraries (e.g., finding albums that are mixed formats or purely one format).

It scans one or more root folders and prints a tree of directories containing files with selected extensions. It can group results by file counts or filter the output to directories that only contain a single extension.

**Usage:**

```bash
extdirisolate [flags]
```

**Flags:**
- `-exts`: Comma-separated list of file extensions to search (default ".flac,.mp3").
- `-roots`: Comma-separated list of root directories to scan (default ".").
- `-case-insensitive`: Make file extension matching case-insensitive.
- `-group`: Group output folders by total file counts.
- `-filter-ext`: Only list folders containing this single extension.
- `-print-parent`: Print parent filtering logic (default true).

**Example:**

To find all directories in `/music` that contain `.flac` or `.mp3` files:

```bash
extdirisolate -exts .flac,.mp3 -roots /music
```

To find directories that *only* contain `.flac` files:

```bash
extdirisolate -exts .flac,.mp3 -roots /music -filter-ext .flac
```

### dirquery

`dirquery` is a search tool that lists directories containing at least one file matching a regular expression and/or a specific file extension. This is useful for finding projects or modules that contain specific files (e.g., finding all directories containing a `main.go` file).

**Usage:**

```bash
dirquery [flags] [directories...]
```

**Flags:**
- `-ext`: File extension to match (including the leading dot, e.g., `.go`).
- `-pattern`: Regular expression applied to file names.
- `-max-depth`: Maximum directory depth to search (-1 for unlimited).

**Example:**

To find all directories containing `.go` files in the current directory:

```bash
dirquery -ext .go .
```

To find all directories containing files matching "test" in their name within `./src`:

```bash
dirquery -pattern test ./src
```
