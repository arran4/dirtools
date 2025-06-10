# Dirtools

Dirtools provides utilities for working with directory trees. The commands
allow querying and isolating directories based on file extensions.

## Building

This project requires Go 1.23 or newer. Build each command with `go build`:

```bash
go build ./cmd/extdirisolate
go build ./cmd/dirquery
```

## Running

### extdirisolate

The `extdirisolate` command analyzes directories and prints folder
structures based on file extensions.

```bash
./extdirisolate -h
```

See the command help output for all available flags.

### dirquery

`dirquery` is planned for future releases. When implemented, build and run
it the same way:

```bash
./dirquery -h
```

## Contributing

Issues and pull requests are welcome. Please open them on GitHub to discuss
changes or report problems.

## License

Dirtools is licensed under the terms of the GNU General Public License
version 3 (GPLv3). See the [LICENSE](LICENSE) file for full details.
