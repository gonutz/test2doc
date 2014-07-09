package main

import (
	t2d "code.google.com/p/test2doc/test2doc"
	"fmt"
	"os"
	"path/filepath"
)

type NameChopper interface {
	Chop(name string) []string
}

type FileWalker interface {
	Walk(path string, f os.FileInfo, err error) error
}

type TestExtractor interface {
	ExtractFromFile(path string) ([]string, error)
}

type DocFormatter interface {
	Append(testFilePath string, sentences [][]string)
	Format() string
}

type DocCreator struct {
	chopper       NameChopper
	filter        FileWalker
	testExtractor TestExtractor
	formatter     DocFormatter
}

func main() {
	os.Args = append(os.Args, "C:/Users/Lars/Documents/gocode/src/code.google.com/test2doc")
	if len(os.Args) != 2 {
		fmt.Println("Please specify the root directory of your code.")
		return
	}

	folder := os.Args[1]
	fileFilter := t2d.TestFileFilter{}
	testExtractor := t2d.NewTestNameExtractor()
	nameChopper := t2d.CamelCaseChopper{}
	formatter := t2d.NewListDocFormatter(folder)

	filepath.Walk(folder, fileFilter.Walk)
	for _, path := range fileFilter.Files {
		appendTests(path, testExtractor, nameChopper, formatter)
	}
	fmt.Print(formatter.Format())
}

func appendTests(fileName string, extractor TestExtractor, chopper NameChopper, formatter DocFormatter) {
	sentences := make([][]string, 0, 50)
	names, err := extractor.ExtractFromFile(fileName)
	if err != nil {
		fmt.Println("error in ", fileName, ":", err)
	}
	for _, name := range names {
		sentences = append(sentences, chopper.Chop(name))
	}
	formatter.Append(fileName, sentences)
}
