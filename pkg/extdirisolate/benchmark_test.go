package extdirisolate

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func setupBenchmarkDir(b *testing.B, depth int, dirsPerLevel int, filesPerDir int) string {
	rootDir := b.TempDir()

	var createDir func(path string, currentDepth int)
	createDir = func(path string, currentDepth int) {
		if currentDepth >= depth {
			return
		}

		for i := 0; i < dirsPerLevel; i++ {
			dirName := filepath.Join(path, fmt.Sprintf("dir_%d", i))
			if err := os.Mkdir(dirName, 0755); err != nil {
				b.Fatal(err)
			}
			for j := 0; j < filesPerDir; j++ {
				fileName := filepath.Join(dirName, fmt.Sprintf("file_%d.txt", j))
				if err := os.WriteFile(fileName, []byte("content"), 0644); err != nil {
					b.Fatal(err)
				}
			}
			createDir(dirName, currentDepth+1)
		}
	}

	createDir(rootDir, 0)
	return rootDir
}

func BenchmarkBuildFolderCounts(b *testing.B) {
	// Setup a directory structure: depth 3, 5 dirs per level, 10 files per dir
	// This results in significant file system traversal
	rootDir := setupBenchmarkDir(b, 3, 5, 10)

	extensions := []string{".txt", ".go", ".md"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := BuildFolderCounts([]string{rootDir}, extensions, false)
		if err != nil {
			b.Fatalf("BuildFolderCounts failed: %v", err)
		}
	}
}
