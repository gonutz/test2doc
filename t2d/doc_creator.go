package t2d

import (
	"os"
	"path/filepath"
)

type DocCreator struct {
	collector FileCollector
	filter    FileFilter
	extractor TestExtractor
	chopper   NameChopper
	formatter DocFormatter
}

func NewDocCreator(
	fc FileCollector,
	ff FileFilter,
	te TestExtractor,
	c NameChopper,
	df DocFormatter) DocCreator {
	return DocCreator{fc, ff, te, c, df}
}

func (doc DocCreator) CreateDocFromFolder(path string) (string, error) {
	err := doc.collector.Collect(path)
	if err != nil {
		err = doc.tryCollectingRelativeToGOPATH(path)
		if err != nil {
			return "", err
		}
	}
	for _, path := range doc.collector.Paths() {
		doc.extractTestsIfTestFile(path)
	}
	return doc.formatter.Format(), nil
}

func (doc DocCreator) tryCollectingRelativeToGOPATH(path string) error {
	withGoPath := filepath.Join(os.Getenv("GOPATH"), "src", path)
	withGoPath = filepath.ToSlash(withGoPath)
	return doc.collector.Collect(withGoPath)
}

func (doc DocCreator) extractTestsIfTestFile(path string) {
	if doc.filter.IsValid(path) {
		testNames, err := doc.extractor.ExtractTestsFromFile(path)
		if err == nil {
			doc.appendChoppedTestNames(testNames, path)
		}
	}
}

func (doc DocCreator) appendChoppedTestNames(testNames []string, path string) {
	sentences := make([][]string, 0, 20)
	for _, name := range testNames {
		sentences = append(sentences, doc.chopper.Chop(name))
	}
	doc.formatter.Append(path, sentences)
}
