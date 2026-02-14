package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/arran4/dirtools/pkg/extdirisolate"
)

// IsolateCmd is a subcommand `extdirisolate`
// Isolates directories based on file extensions.
//
// Flags:
//
//	exts:            --exts            (default: ".flac,.mp3") List of file extensions (comma separated)
//	roots:           --roots           (default: ".")          List of root directories (comma separated)
//	caseInsensitive: --case-insensitive                        Make file extension matching ignore case
//	groupOutput:     --group                                   Group output folders by total file counts
//	filterExt:       --filter-ext                              Only list folders containing this single extension
//	printParent:     --print-parent    (default: true)         Print parent filtering logic
func IsolateCmd(exts string, roots string, caseInsensitive bool, groupOutput bool, filterExt string, printParent bool) error {
	extensions := strings.Split(exts, ",")
	if caseInsensitive {
		for i, ext := range extensions {
			extensions[i] = strings.ToLower(ext)
		}
	}

	rootDirs := strings.Split(roots, ",")

	folderCounts, err := extdirisolate.BuildFolderCounts(rootDirs, extensions, caseInsensitive)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	for _, rootDir := range rootDirs {
		if fc, ok := folderCounts[rootDir]; ok {
			fc.Print(filterExt, groupOutput, 0, printParent)
		} else {
            // Should not happen
            fmt.Fprintf(os.Stderr, "Warning: root dir %s not found in results\n", rootDir)
        }
	}
	return nil
}
