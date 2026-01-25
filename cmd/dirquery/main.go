package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arran4/dirtools/pkg/dirquery"
)

func main() {
	var pattern string
	var ext string
	var maxDepth int
	flag.StringVar(&pattern, "pattern", "", "regular expression for filenames")
	flag.StringVar(&ext, "ext", "", "file extension to match (with dot)")
	flag.IntVar(&maxDepth, "max-depth", -1, "maximum directory depth to search")
	flag.Parse()

	dirs := flag.Args()

	err := dirquery.Search(dirs, pattern, ext, maxDepth, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
