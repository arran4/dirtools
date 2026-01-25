package extdirisolate

import (
	"os"
	"path/filepath"
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

func TestBuildFolderCountsTraversal(t *testing.T) {
	root := t.TempDir()
	createTestFile(t, filepath.Join(root, "a", "file1.mp3"))
	createTestFile(t, filepath.Join(root, "a", "file2.flac"))
	createTestFile(t, filepath.Join(root, "b", "file3.mp3"))
	createTestFile(t, filepath.Join(root, "b", "c", "file4.flac"))

	counts, err := BuildFolderCounts([]string{root}, []string{".mp3", ".flac"}, false)
	if err != nil {
		t.Fatalf("BuildFolderCounts error: %v", err)
	}

	rootFC := counts[root]
	if rootFC == nil {
		t.Fatalf("root folder missing")
	}
	if got := rootFC.TotalFiles; got != 4 {
		t.Errorf("root total files = %d, want 4", got)
	}
	if got := rootFC.Counts[".mp3"]; got != 2 {
		t.Errorf("root mp3 count = %d, want 2", got)
	}
	if got := rootFC.Counts[".flac"]; got != 2 {
		t.Errorf("root flac count = %d, want 2", got)
	}

	// check child relationship for nested directory
	bcPath := filepath.Join(root, "b", "c")
	bc := counts[bcPath]
	b := counts[filepath.Dir(bcPath)]
	if bc == nil || b == nil {
		t.Fatalf("missing folder counts for nested directories")
	}
	if bc.Parent != b {
		t.Errorf("expected %s parent to be %s", bcPath, filepath.Join(root, "b"))
	}
	if got := bc.Counts[".flac"]; got != 1 {
		t.Errorf("b/c flac count = %d, want 1", got)
	}
}
