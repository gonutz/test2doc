package t2d

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilesNotEndingIn_UnderscoreTestDotGo_AreInvalid(t *testing.T) {
	filter := NewTestFileFilter()
	invalid := []string{
		"some/path/not_this.go",
		"some/path/no_test.ogg",
		"dont/take/thistest.go"}
	for _, path := range invalid {
		if filter.IsValid(path) {
			t.Error(path, "was valid but should not be")
		}
	}
}

func TestFoldersAreNeverValidTestFiles(t *testing.T) {
	filter := NewTestFileFilter()
	dir := createFolder("folder_test.go")
	defer deleteFile(dir)
	if filter.IsValid(dir) {
		t.Error("folder was valid")
	}
}

func TestRegularFilesEndingIn_UnderscoreTestDotGo_AreValid(t *testing.T) {
	filter := NewTestFileFilter()
	f := createFileAt(filepath.Join(os.TempDir(), "some_test.go"))
	defer deleteFile(f)
	if !filter.IsValid(f) {
		t.Error(f, "was invalid")
	}
}

func TestFilesThatDontExistsAreInvalid(t *testing.T) {
	filter := NewTestFileFilter()
	if filter.IsValid("invalid///path/ending_in_test.go") {
		t.Error("non-existing file was valid")
	}
}

func TestIfFileNameIsExactly_UnderscoreTestDotGo_ItIsInvalid(t *testing.T) {
	filter := NewTestFileFilter()
	path := createFileAt(filepath.Join(os.TempDir(), "_test.go"))
	if filter.IsValid(path) {
		t.Error("_test.go must be invalid")
	}
}
