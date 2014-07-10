package t2d

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
		return "", err
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
