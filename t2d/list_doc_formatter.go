package t2d

import (
	"fmt"
	"path"
	"strings"
)

type ListDocFormatter struct {
	doc  string
	root string
}

func NewListDocFormatter(cutPathPrefix string) *ListDocFormatter {
	slashPath := forwardSlashesOnly(cutPathPrefix)
	return &ListDocFormatter{root: slashPath}
}

func forwardSlashesOnly(path string) string {
	return strings.Replace(path, `\`, "/", -1)
}

func (f *ListDocFormatter) Format() string {
	return f.doc
}

func (f *ListDocFormatter) Append(testFilePath string, sentences [][]string) {
	f.insertNewLine()
	f.appendUnitName(testFilePath)
	for _, sentence := range sentences {
		f.appendSentence(sentence)
	}
}

func (f *ListDocFormatter) insertNewLine() {
	if len(f.doc) > 0 {
		f.doc += fmt.Sprintln()
	}
}

func (f *ListDocFormatter) appendUnitName(testFilePath string) {
	slashes := forwardSlashesOnly(testFilePath)
	withoutRoot := cutRootPrefix(slashes, f.root)
	unitName := cutTestSuffix(withoutRoot)

	f.doc += fmt.Sprintln(unitName + ":")
}

func cutTestSuffix(path string) string {
	return strings.TrimSuffix(path, "_test.go")
}

func cutRootPrefix(file, root string) string {
	withoutRoot := strings.TrimPrefix(file, root)
	if len(withoutRoot) == 0 {
		_, withoutRoot = path.Split(file)
	}
	withoutRoot = strings.TrimPrefix(withoutRoot, "/")
	return withoutRoot
}

func (f *ListDocFormatter) appendSentence(sentence []string) {
	f.doc += fmt.Sprintln("    - " + strings.Join(sentence, " "))
}
