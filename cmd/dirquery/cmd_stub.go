package main

import (
	"fmt"
	"os"
    "strconv"

	"github.com/arran4/dirtools/pkg/dirquery"
)

// SearchCmd is a subcommand `dirquery`
// Searches directories for folders containing files matching the given pattern and extension.
//
// Flags:
//
//	pattern:  --pattern -p Regular expression for filenames
//	ext:      --ext -e     File extension to match (with dot)
//	maxDepthStr: --max-depth -d Maximum directory depth to search (defaults to minus 1)
//	dirs:     ...          Directories to search
func SearchCmd(pattern string, ext string, maxDepthStr string, dirs ...string) error {
    maxDepth := -1
    if maxDepthStr != "" {
        var err error
        maxDepth, err = strconv.Atoi(maxDepthStr)
        if err != nil {
            return fmt.Errorf("invalid max-depth: %v", err)
        }
    }

	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	err := dirquery.Search(dirs, pattern, ext, maxDepth, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}
