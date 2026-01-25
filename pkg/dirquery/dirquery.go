package dirquery

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Search scans directories for folders containing files matching the given pattern and extension.
// matches are written to out. Errors encountered during traversal are written to errOut.
func Search(dirs []string, pattern string, ext string, maxDepth int, out io.Writer, errOut io.Writer) error {
	var re *regexp.Regexp
	var err error
	if pattern != "" {
		re, err = regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("invalid pattern: %v", err)
		}
	}

	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	for _, dir := range dirs {
		rootDepth := strings.Count(filepath.Clean(dir), string(os.PathSeparator))
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				if errOut != nil {
					fmt.Fprintln(errOut, err)
				}
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
				fmt.Fprintln(out, path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
