package test2doc

import (
	"os"
	"strings"
)

type TestFileFilter struct {
	Files []string
}

func (t *TestFileFilter) Walk(path string, f os.FileInfo, err error) error {
	if err == nil && strings.HasSuffix(path, "_test.go") && !f.IsDir() {
		t.Files = append(t.Files, path)
	}
	return nil
}
