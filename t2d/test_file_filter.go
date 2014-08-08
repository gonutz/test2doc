package t2d

import (
	"os"
	"path/filepath"
	"strings"
)

type TestFileFilter struct {
	Files []string
}

func NewTestFileFilter() TestFileFilter {
	return TestFileFilter{}
}

func (_ TestFileFilter) IsValid(fullPath string) bool {
	if isFileAndNotDirectory(fullPath) {
		_, file := filepath.Split(fullPath)
		return strings.HasSuffix(file, "_test.go") && file != "_test.go"
	}
	return false
}

func isFileAndNotDirectory(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir()
	}
	return false
}
