package dirquery

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	root := b.TempDir()
	// Create a structure with many directories to exercise the directory path logic
	for i := 0; i < 1000; i++ {
		dir := filepath.Join(root, fmt.Sprintf("dir_%d", i))
		if err := os.MkdirAll(dir, 0755); err != nil {
			b.Fatalf("failed to create dir: %v", err)
		}
        // One nested directory
        nested := filepath.Join(dir, "nested")
        if err := os.MkdirAll(nested, 0755); err != nil {
             b.Fatalf("failed to create nested dir: %v", err)
        }
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// We use a maxDepth >= 0 to trigger the affected code path
		err := Search([]string{root}, "", "", 10, io.Discard, io.Discard)
		if err != nil {
			b.Fatalf("Search failed: %v", err)
		}
	}
}
