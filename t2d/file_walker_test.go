package t2d

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIfPathDoesNotExistsAnErrorIsReturned(t *testing.T) {
	walker := NewFileWalker()
	err := walker.Collect("invalid///path")
	if err != PathDoesNotExist {
		t.Error("PathDoesNotExist error expected but was", err)
	}
}

func TestIfPathIsAFileOnlyThatFileIsWalked(t *testing.T) {
	walker := NewFileWalker()
	f := createFile("")
	defer deleteFile(f)

	err := walker.Collect(f)

	if err != nil {
		t.Error("No error expected but was", err)
	}
	checkNames(t, walker.Paths(), f)
}

func TestAllPathsInTreeAreWalked(t *testing.T) {
	walker := NewFileWalker()
	root, sub, f1, f2 := createTestFolderTree()
	defer deleteTestFolderTree(root, sub, f1, f2)

	err := walker.Collect(root)

	if err != nil {
		t.Error("No error expected but was", err)
	}
	checkNames(t, walker.Paths(), root, f1, sub, f2)
}

func createTestFolderTree() (root, sub, f1, f2 string) {
	root = createFolder("root")
	sub = createFolder(filepath.Join("root", "sub"))
	f1 = createFileAt(filepath.Join(root, "f1"))
	f2 = createFileAt(filepath.Join(sub, "f2"))
	return
}

func createFolder(name string) string {
	path := filepath.Join(os.TempDir(), name)
	os.MkdirAll(path, os.ModePerm)
	return path
}

func deleteTestFolderTree(root, sub, f1, f2 string) {
	deleteFile(f1)
	deleteFile(f2)
	deleteFile(sub)
	deleteFile(root)
}

func createFileAt(path string) string {
	f, _ := os.Create(path)
	f.Close()
	return path
}
