package t2d

type FileCollector interface {
	Collect(path string) error
	Paths() []string
}

type FileFilter interface {
	IsValid(paths string) bool
}

type TestExtractor interface {
	ExtractTestsFromFile(path string) ([]string, error)
}

type NameChopper interface {
	Chop(name string) []string
}

type DocFormatter interface {
	Append(testFilePath string, sentences [][]string)
	Format() string
}
