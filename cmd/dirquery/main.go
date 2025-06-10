package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var pattern string
	var ext string
	var maxDepth int
	flag.StringVar(&pattern, "pattern", "", "regular expression for filenames")
	flag.StringVar(&ext, "ext", "", "file extension to match (with dot)")
	flag.IntVar(&maxDepth, "max-depth", -1, "maximum directory depth to search")
	flag.Parse()

	var re *regexp.Regexp
	var err error
	if pattern != "" {
		re, err = regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid pattern: %v\n", err)
			os.Exit(1)
		}
	}

	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	for _, dir := range dirs {
		rootDepth := strings.Count(filepath.Clean(dir), string(os.PathSeparator))
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			if !d.IsDir() {
				return nil
			}
			if maxDepth >= 0 {
				depth := strings.Count(filepath.Clean(path), string(os.PathSeparator)) - rootDepth
				if depth > maxDepth {
					return filepath.SkipDir
				}
			}
			entries, err := os.ReadDir(path)
			if err != nil {
				return nil
			}
			matched := false
			for _, e := range entries {
				if e.IsDir() {
					continue
				}
				name := e.Name()
				if ext != "" && filepath.Ext(name) != ext {
					continue
				}
				if re != nil && !re.MatchString(name) {
					continue
				}
				matched = true
				break
			}
			if matched {
				fmt.Println(path)
			}
			return nil
		})
	}
}
