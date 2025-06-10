package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FolderCounts struct {
	Counts     map[string]int
	Children   []*FolderCounts
	Parent     *FolderCounts
	Name       string
	TotalFiles int
}

// Add updates the counts and propagates them up to parent folders
func (fc *FolderCounts) Add(extension string, count int) {
	fc.Counts[extension] += count
	fc.TotalFiles += count
	if fc.Parent != nil {
		fc.Parent.Add(extension, count)
	}
}

// Mixed determines if a folder has multiple file types
func (fc *FolderCounts) Mixed() bool {
	unique := 0
	for _, count := range fc.Counts {
		if count > 0 {
			unique++
		}
	}
	return unique > 1
}

// Print outputs the folder structure based on conditions
func (fc *FolderCounts) Print(onlyExt string, group bool, level int, printParent bool) {
	if printParent && fc.Parent != nil && fc.Parent.Mixed() && !fc.Mixed() {
		fmt.Printf("First non-mixed folder: %s\n", fc.Name)
		return
	}

	if onlyExt != "" {
		if fc.Counts[onlyExt] > 0 && len(fc.Counts) == 1 {
			fmt.Printf("%s%s\n", strings.Repeat("  ", level), fc.Name)
		}
	} else {
		if group {
			fmt.Printf("%s%s (%d files)\n", strings.Repeat("  ", level), fc.Name, fc.TotalFiles)
		} else {
			fmt.Printf("%s%s\n", strings.Repeat("  ", level), fc.Name)
		}
	}

	for _, child := range fc.Children {
		child.Print(onlyExt, group, level+1, printParent)
	}
}

// BuildFolderCounts walks the provided root directories and returns a map of
// FolderCounts keyed by path. The logic mirrors the traversal used by the main
// command so it can be reused in tests.
func BuildFolderCounts(rootDirs []string, extensions []string, caseInsensitive bool) (map[string]*FolderCounts, error) {
	folderCounts := make(map[string]*FolderCounts)

	for _, rootDir := range rootDirs {
		folderCounts[rootDir] = &FolderCounts{Name: rootDir, Counts: make(map[string]int)}
		err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				if _, exists := folderCounts[path]; !exists {
					folderCounts[path] = &FolderCounts{Name: path, Counts: make(map[string]int)}
				}

				parentPath := filepath.Dir(path)
				if parentPath != path { // Avoid linking root to itself
					parentFolder, parentExists := folderCounts[parentPath]
					if !parentExists {
						parentFolder = &FolderCounts{Name: parentPath, Counts: make(map[string]int)}
						folderCounts[parentPath] = parentFolder
					}
					parentFolder.Children = append(parentFolder.Children, folderCounts[path])
					folderCounts[path].Parent = parentFolder
				}
				return nil
			}

			folder := filepath.Dir(path)
			currentFolder, exists := folderCounts[folder]
			if !exists {
				currentFolder = &FolderCounts{Name: folder, Counts: make(map[string]int)}
				folderCounts[folder] = currentFolder
			}

			ext := filepath.Ext(path)
			if caseInsensitive {
				ext = strings.ToLower(ext)
			}
			for _, validExt := range extensions {
				if ext == validExt {
					currentFolder.Add(validExt, 1)
					break
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return folderCounts, nil
}

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

	folderCounts, err := BuildFolderCounts(rootDirs, extensions, caseInsensitive)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print the folder structure
	for _, rootDir := range rootDirs {
		folderCounts[rootDir].Print(filterExt, groupOutput, 0, printParent)
	}
}
