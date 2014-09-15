package t2d

import (
	"errors"
	"os"
	"testing"
)

func TestIfFileCollectorFails_PathWith_GopathSlashSrc_IsTried(t *testing.T) {
	os.Setenv("GOPATH", "go/path")
	spy := &firstFailSpyCollector{}
	creator := NewDocCreator(
		spy,
		yesFilter{},
		dummyExtractor{},
		dummyChopper{},
		dummyFormatter{})
	creator.CreateDocFromFolder("folder")
	if spy.secondFolder != "go/path/src/folder" {
		t.Error("GOPATH was not tried after failure but was", spy.secondFolder)
	}
}

func TestIfFileCollectorFailsEvenWithGoPath_ItsErrorIsReturned(t *testing.T) {
	expected := errors.New("test")
	creator := NewDocCreator(
		failingCollector{expected},
		yesFilter{},
		dummyExtractor{},
		dummyChopper{},
		dummyFormatter{})
	doc, err := creator.CreateDocFromFolder("")
	if err != expected {
		t.Error("error expected but was", err)
	}
	if doc != "" {
		t.Error("on error the doc must be empty string")
	}
}

func TestCollectedFilesAreFiltered(t *testing.T) {
	spyFilter := &spyNameFilter{}
	creator := NewDocCreator(
		stubCollector{[]string{"1", "2", "3"}},
		spyFilter,
		dummyExtractor{},
		dummyChopper{},
		dummyFormatter{})
	creator.CreateDocFromFolder("")
	checkNames(t, spyFilter.files, "1", "2", "3")
}

func TestFilteredFilesAreGivenToExtractor(t *testing.T) {
	spyExtractor := &spyTestExtractor{}
	creator := NewDocCreator(
		stubCollector{[]string{"file1", "file2", "file3", "file4", "file5"}},
		&alternatingFilter{},
		spyExtractor,
		dummyChopper{},
		dummyFormatter{})
	creator.CreateDocFromFolder("")
	checkNames(t, spyExtractor.files, "file1", "file3", "file5")
}

func TestEachNameIsChopped(t *testing.T) {
	spyChopper := &spyNameChopper{}
	creator := NewDocCreator(
		stubCollector{[]string{"file"}},
		yesFilter{},
		stubExtractor{[]string{"test1", "test2", "test3"}},
		spyChopper,
		dummyFormatter{})
	creator.CreateDocFromFolder("")
	checkNames(t, spyChopper.names, "test1", "test2", "test3")
}

func TestIfExtractingTestFails_NoNameIsChoppedForTheFile(t *testing.T) {
	spyChopper := &spyNameChopper{}
	creator := NewDocCreator(
		stubCollector{[]string{"file"}},
		yesFilter{},
		failingExtractor{errors.New("some error")},
		spyChopper,
		dummyFormatter{})
	creator.CreateDocFromFolder("")
	checkNames(t, spyChopper.names)
}

func TestChoppedNamesArePassedToTheFormatter(t *testing.T) {
	spyFormatter := &spyDocFormatter{}
	creator := NewDocCreator(
		stubCollector{[]string{"file"}},
		yesFilter{},
		stubExtractor{[]string{"test1", "test2"}},
		stubChopper{[]string{"word1", "word2", "word3"}},
		spyFormatter)
	creator.CreateDocFromFolder("")
	checkNames(t, spyFormatter.paths, "file")
	checkNames(t, spyFormatter.paragraphs[0][0], "word1", "word2", "word3")
	checkNames(t, spyFormatter.paragraphs[0][1], "word1", "word2", "word3")
}

func TestAtTheEndTheFormattedDocStringIsReturned(t *testing.T) {
	creator := NewDocCreator(
		stubCollector{[]string{"file"}},
		yesFilter{},
		dummyExtractor{},
		dummyChopper{},
		stubFormatter{"this is the final doc"})
	doc, _ := creator.CreateDocFromFolder("")
	if doc != "this is the final doc" {
		t.Error("wrong doc string:", doc)
	}
}

type firstFailSpyCollector struct {
	secondFolder string
	tries        int
}

func (c *firstFailSpyCollector) Collect(path string) error {
	c.tries++
	if c.tries == 2 {
		c.secondFolder = path
		return nil
	}
	return errors.New("first try fails")
}
func (_ *firstFailSpyCollector) Paths() []string { return []string{} }

type failingCollector struct{ err error }

func (c failingCollector) Collect(path string) error { return c.err }
func (_ failingCollector) Paths() []string           { return []string{} }

type stubCollector struct{ files []string }

func (_ stubCollector) Collect(path string) error { return nil }
func (c stubCollector) Paths() []string           { return c.files }

type yesFilter struct{}

func (_ yesFilter) IsValid(_ string) bool { return true }

type spyNameFilter struct{ files []string }

func (spy *spyNameFilter) IsValid(path string) bool {
	spy.files = append(spy.files, path)
	return false
}

type alternatingFilter struct{ answer bool }

func (f *alternatingFilter) IsValid(path string) bool {
	f.answer = !f.answer
	return f.answer
}

type dummyExtractor struct{}

func (_ dummyExtractor) ExtractTestsFromFile(_ string) ([]string, error) { return []string{}, nil }

type spyTestExtractor struct{ files []string }

func (spy *spyTestExtractor) ExtractTestsFromFile(path string) ([]string, error) {
	spy.files = append(spy.files, path)
	return []string{}, nil
}

type stubExtractor struct{ tests []string }

func (stub stubExtractor) ExtractTestsFromFile(_ string) ([]string, error) { return stub.tests, nil }

type failingExtractor struct{ err error }

func (e failingExtractor) ExtractTestsFromFile(_ string) ([]string, error) {
	return []string{"a"}, e.err
}

type dummyChopper struct{}

func (_ dummyChopper) Chop(_ string) []string { return []string{} }

type spyNameChopper struct{ names []string }

func (spy *spyNameChopper) Chop(name string) []string {
	spy.names = append(spy.names, name)
	return []string{}
}

type stubChopper struct{ words []string }

func (stub stubChopper) Chop(_ string) []string { return stub.words }

type dummyFormatter struct{}

func (_ dummyFormatter) Append(_ string, _ [][]string) {}
func (_ dummyFormatter) Format() string                { return "" }

type spyDocFormatter struct {
	paths      []string
	paragraphs [][][]string
}

func (spy *spyDocFormatter) Append(testFilePath string, sentences [][]string) {
	spy.paths = append(spy.paths, testFilePath)
	spy.paragraphs = append(spy.paragraphs, sentences)
}
func (spy *spyDocFormatter) Format() string { return "" }

type stubFormatter struct{ doc string }

func (_ stubFormatter) Append(_ string, _ [][]string) {}
func (stub stubFormatter) Format() string             { return stub.doc }
