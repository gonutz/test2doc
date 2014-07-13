package t2d

import (
	"bytes"
	"fmt"
	"path"
	"strings"
)

type ListDocFormatter struct {
	doc     bytes.Buffer
	root    string
	newLine string
}

func NewListDocFormatter(cutPathPrefix string) *ListDocFormatter {
	slashPath := forwardSlashesOnly(cutPathPrefix)
	return &ListDocFormatter{root: slashPath, newLine: fmt.Sprintln()}
}

func forwardSlashesOnly(path string) string {
	return strings.Replace(path, `\`, "/", -1)
}

func (f *ListDocFormatter) Format() string {
	return f.doc.String()
}

func (f *ListDocFormatter) Append(testFilePath string, sentences [][]string) {
	f.insertNewLine()
	f.appendUnitName(testFilePath)
	for _, sentence := range sentences {
		f.appendSentence(sentence)
	}
}

func (f *ListDocFormatter) insertNewLine() {
	if f.doc.Len() > 0 {
		f.doc.WriteString(f.newLine)
	}
}

func (f *ListDocFormatter) appendUnitName(testFilePath string) {
	slashes := forwardSlashesOnly(testFilePath)
	withoutRoot := cutRootPrefix(slashes, f.root)
	unitName := cutTestSuffix(withoutRoot)

	f.doc.WriteString(unitName)
	f.doc.WriteString(":")
	f.doc.WriteString(f.newLine)
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
	if len(sentence) > 0 {
		f.doc.WriteString("    - ")
		f.doc.WriteString(strings.Join(sentence, " "))
		f.doc.WriteString(f.newLine)
	}
}
