package t2d

import (
	"os"
	"strings"
)

type TestFileFilter struct {
	Files []string
}

func NewTestFileFilter() TestFileFilter {
	return TestFileFilter{}
}

func (_ TestFileFilter) IsValid(path string) bool {
	if strings.HasSuffix(path, "_test.go") {
		return isFileAndNotDirectory(path)
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
