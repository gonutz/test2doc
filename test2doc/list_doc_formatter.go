package test2doc

import (
	"fmt"
	"strings"
)

type ListDocFormatter struct {
	doc  string
	path string
}

func NewListDocFormatter(cutPathPrefix string) *ListDocFormatter {
	slashPath := strings.Replace(cutPathPrefix, `\`, "/", -1)
	return &ListDocFormatter{path: slashPath}
}

func (f *ListDocFormatter) Format() string {
	return f.doc
}

func (f *ListDocFormatter) Append(testFilePath string, sentences [][]string) {
	if len(f.doc) > 0 {
		f.doc += fmt.Sprintln()
	}
	slashes := strings.Replace(testFilePath, `\`, "/", -1)
	pathCut := strings.TrimPrefix(slashes, f.path)
	pathCut = strings.TrimPrefix(pathCut, "/")
	unitName := strings.TrimSuffix(pathCut, "_test.go")
	f.doc += fmt.Sprintln(unitName + ":")
	for _, sentence := range sentences {
		f.doc += fmt.Sprintln("    - " + strings.Join(sentence, " "))
	}
}
