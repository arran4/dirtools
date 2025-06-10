# Dirtools

Dirtools is a small collection of command line utilities for working with
directories.

## Commands

### extdirisolate

Scans one or more root folders and prints a tree of directories containing files
with selected extensions.  It can group results by file counts or filter the
output to directories that only contain a single extension.

Example:

```bash
extdirisolate -exts .flac,.mp3 -roots /music
```

### dirquery

Lists directories that contain at least one file matching a regular expression
and/or file extension.  Multiple roots can be provided as positional arguments.

Common flags:

- `-ext` sets the file extension to match (including the leading dot).
- `-pattern` regular expression applied to file names.
- `-max-depth` controls how deep the search may recurse (-1 for unlimited).

Example:

```bash
dirquery -ext .go -pattern test ./src
```
