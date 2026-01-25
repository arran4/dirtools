package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/arran4/dirtools/pkg/extdirisolate"
)

func main() {
	// Command-line arguments
	var exts string
	var roots string
	var caseInsensitive bool
	var groupOutput bool
	var filterExt string
	var printParent bool

	flag.StringVar(&exts, "exts", ".flac,.mp3", "Comma-separated list of file extensions to search")
	flag.StringVar(&roots, "roots", ".", "Comma-separated list of root directories")
	flag.BoolVar(&caseInsensitive, "case-insensitive", false, "Make file extension matching case-insensitive")
	flag.BoolVar(&groupOutput, "group", false, "Group output folders by total file counts")
	flag.StringVar(&filterExt, "filter-ext", "", "Only list folders containing this single extension")
	flag.BoolVar(&printParent, "print-parent", true, "Print parent filtering logic")
	flag.Parse()

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
		os.Exit(1)
	}

	// Print the folder structure
	for _, rootDir := range rootDirs {
		folderCounts[rootDir].Print(filterExt, groupOutput, 0, printParent)
	}
}
