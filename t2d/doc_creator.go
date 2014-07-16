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

func NewDocCreator(fc FileCollector, ff FileFilter, te TestExtractor, c NameChopper, df DocFormatter) DocCreator {
	return DocCreator{fc, ff, te, c, df}
}

func (doc DocCreator) CreateDocFromFolder(path string) (string, error) {
	err := doc.collector.Collect(path)
	if err != nil {
		withGoPath := filepath.Join(os.Getenv("GOPATH"), "src", path)
		withGoPath = filepath.ToSlash(withGoPath)
		err = doc.collector.Collect(withGoPath)
		if err != nil {
			return "", err
		}
	}
	for _, path := range doc.collector.Paths() {
		if doc.filter.IsValid(path) {
			testNames, err := doc.extractor.ExtractTestsFromFile(path)
			if err == nil {
				sentences := make([][]string, 0, 20)
				for _, name := range testNames {
					sentences = append(sentences, doc.chopper.Chop(name))
				}
				doc.formatter.Append(path, sentences)
			}
		}
	}
	return doc.formatter.Format(), nil
}
