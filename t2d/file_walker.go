package t2d

import (
	"errors"
	"os"
	"path/filepath"
)

var PathDoesNotExist = errors.New("The given path does not exist.")

type FileWalker struct {
	names []string
}

func NewFileWalker() *FileWalker {
	return &FileWalker{make([]string, 0, 100)}
}

// Collect collects all files and folders under path including path itself
func (walker *FileWalker) Collect(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return PathDoesNotExist
	}
	filepath.Walk(path, walker.walkFunction)
	return nil
}

func (walker *FileWalker) walkFunction(path string, f os.FileInfo, err error) error {
	walker.names = append(walker.names, path)
	return nil
}

func (walker *FileWalker) Paths() []string {
	return walker.names
}
