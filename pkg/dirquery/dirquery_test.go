package dirquery

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTestFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte("test"), 0o644); err != nil {
		t.Fatalf("writefile: %v", err)
	}
}

func TestSearch(t *testing.T) {
	root := t.TempDir()
	createTestFile(t, filepath.Join(root, "src", "main.go"))
	createTestFile(t, filepath.Join(root, "src", "utils.go"))
	createTestFile(t, filepath.Join(root, "src", "README.md"))
	createTestFile(t, filepath.Join(root, "test", "main_test.go"))
	createTestFile(t, filepath.Join(root, "vendor", "lib", "lib.go"))

	tests := []struct {
		name     string
		pattern  string
		ext      string
		maxDepth int
		want     []string
	}{
		{
			name:     "Find all go files",
			pattern:  "",
			ext:      ".go",
			maxDepth: -1,
			want: []string{
				filepath.Join(root, "src"),
				filepath.Join(root, "test"),
				filepath.Join(root, "vendor", "lib"),
			},
		},
		{
			name:     "Find main files by regex",
			pattern:  "main",
			ext:      "",
			maxDepth: -1,
			want: []string{
				filepath.Join(root, "src"),
				filepath.Join(root, "test"),
			},
		},
		{
			name:     "Find files with depth limit",
			pattern:  "",
			ext:      ".go",
			maxDepth: 1, // src is depth 1. vendor/lib is depth 2.
			want: []string{
				filepath.Join(root, "src"),
				filepath.Join(root, "test"),
			},
		},
		{
			name:     "Find nothing",
			pattern:  "nomatch",
			ext:      "",
			maxDepth: -1,
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			var errOut bytes.Buffer
			err := Search([]string{root}, tt.pattern, tt.ext, tt.maxDepth, &out, &errOut)
			if err != nil {
				t.Fatalf("Search error: %v", err)
			}
			got := strings.Split(strings.TrimSpace(out.String()), "\n")
			if len(got) == 1 && got[0] == "" {
				got = []string{}
			}

			gotMap := make(map[string]bool)
			for _, g := range got {
				gotMap[g] = true
			}

			if len(got) != len(tt.want) {
				t.Errorf("got %d results, want %d\nGot: %v\nWant: %v", len(got), len(tt.want), got, tt.want)
			}

			for _, w := range tt.want {
				if !gotMap[w] {
					t.Errorf("missing expected path: %s", w)
				}
			}
		})
	}
}
