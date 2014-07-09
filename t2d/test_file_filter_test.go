package t2d

import (
	"os"
	"testing"
	"time"
)

func TestOnlyFilesEndingInUnderscoreTestDotGoAreTaken(t *testing.T) {
	filter := TestFileFilter{}
	filter.Walk("some/path/not_this.go", FileStub{}, nil)
	filter.Walk("some/path/no_test.ogg", FileStub{}, nil)
	filter.Walk("take/this_test.go", FileStub{}, nil)
	filter.Walk("but/not/thistest.go", FileStub{}, nil)
	checkFile(t, filter, "take/this_test.go")
}

func TestDirectoriesAreNeverTaken(t *testing.T) {
	filter := TestFileFilter{}
	filter.Walk("dir_ending_in_test.go", DirStub{}, nil)
	filter.Walk("file_ending_in_test.go", FileStub{}, nil)
	checkFile(t, filter, "file_ending_in_test.go")
}

func TestIfAnErrorOccursInWalkingAFile_TheFileIsNotTaken(t *testing.T) {
	filter := TestFileFilter{}
	filter.Walk("no_error_test.go", FileStub{}, nil)
	filter.Walk("has_error_test.go", FileStub{}, os.ErrPermission)
	checkFile(t, filter, "no_error_test.go")
}

type DirStub struct{}

func (d DirStub) Name() string       { return "" }
func (d DirStub) Size() int64        { return 0 }
func (d DirStub) Mode() os.FileMode  { return 0 }
func (d DirStub) ModTime() time.Time { return time.Time{} }
func (d DirStub) IsDir() bool        { return true }
func (d DirStub) Sys() interface{}   { return nil }

type FileStub struct{}

func (d FileStub) Name() string       { return "" }
func (d FileStub) Size() int64        { return 0 }
func (d FileStub) Mode() os.FileMode  { return 0 }
func (d FileStub) ModTime() time.Time { return time.Time{} }
func (d FileStub) IsDir() bool        { return false }
func (d FileStub) Sys() interface{}   { return nil }

func checkFile(t *testing.T, filter TestFileFilter, path string) {
	if len(filter.Files) != 1 {
		t.Error("1 file expected, but was", filter.Files)
	} else if filter.Files[0] != path {
		t.Error("file was", filter.Files[0])
	}
}
