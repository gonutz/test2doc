package t2d

import "testing"

func TestWithoutAnyTestFileTheOutputIsEmptyString(t *testing.T) {
	formatter := NewListDocFormatter("")
	checkDoc(t, formatter, "")
}

func TestFirstLineIsFileNameWithout_TestDotGo(t *testing.T) {
	formatter := NewListDocFormatter("")
	formatter.Append("some_test.go", [][]string{})
	checkDoc(t, formatter, "some:\n")
}

func TestRootFolderIsTrimmedFromFilePathPrefix(t *testing.T) {
	formatter := NewListDocFormatter("/root/path")
	formatter.Append("/root/path/src/unit_test.go", [][]string{})
	checkDoc(t, formatter, "src/unit:\n")
}

func TestBackslashesInPathAreConvertedToSlashes(t *testing.T) {
	formatter := NewListDocFormatter("")
	formatter.Append(`src\unit_test.go`, [][]string{})
	checkDoc(t, formatter, "src/unit:\n")
}

func TestRootPathCanHaveBackslashAtTheEnd(t *testing.T) {
	formatter := NewListDocFormatter(`\root\path\`)
	formatter.Append(`\root\path\src\unit_test.go`, [][]string{})
	checkDoc(t, formatter, "src/unit:\n")
}

func TestRootPathMayMixSlashesAndBackslashes(t *testing.T) {
	formatter := NewListDocFormatter(`/root\path/`)
	formatter.Append(`\root\path\src\unit_test.go`, [][]string{})
	checkDoc(t, formatter, "src/unit:\n")
}

func TestIfRootPathEqualsPath_TheFileNameIsUsed(t *testing.T) {
	formatter := NewListDocFormatter("path/my_test.go")
	formatter.Append("path/my_test.go", [][]string{})
	checkDoc(t, formatter, "my:\n")
}

func TestEverySentenceIsIndentendOnANewLine(t *testing.T) {
	formatter := NewListDocFormatter("")
	formatter.Append("unit_test.go", [][]string{
		[]string{"a", "b", "c"},
		[]string{"Hello", "World"}})
	checkDoc(t, formatter,
		`unit:
    - a b c
    - Hello World
`)
}

func TestUnitDocsAreSeparatedByABlankLine(t *testing.T) {
	formatter := NewListDocFormatter("")
	formatter.Append("unit_1_test.go", [][]string{})
	formatter.Append("unit_2_test.go", [][]string{})
	checkDoc(t, formatter, "unit_1:\n\nunit_2:\n")
}

func TestEmptySentencesDoNotAppearInTheDoc(t *testing.T) {
	formatter := NewListDocFormatter("")
	emptySentence := []string{}
	formatter.Append("unit_test.go", [][]string{emptySentence})
	checkDoc(t, formatter, "unit:\n")
}

func checkDoc(t *testing.T, formatter *ListDocFormatter, expected string) {
	doc := formatter.Format()
	if doc != expected {
		t.Error(expected, "expected but was", doc)
	}
}
