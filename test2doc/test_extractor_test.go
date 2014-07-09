package test2doc

import (
	"io/ioutil"
	"os"
	"testing"
)

// This test repeats the same operation 100 times, expecting the same output
// each time. Due to concurrency issues, the implementation may give different
// results from run to run so this is a statistical test that may return a false
// positive every now and then.
func TestTestsFromFileAreExtractedInOrderOfAppearance(t *testing.T) {
	file := createFile(`
			package sometest
			import "testing"

			func TestSomeThing1(t *testing.T){}
			func TestSomeThing2(t *testing.T){}
			func TestSomeThing3(t *testing.T){}`)
	defer deleteFile(file)

	for i := 0; i < 100; i++ {
		if !checkExtractsNames(t, file.Name(),
			"TestSomeThing1",
			"TestSomeThing2",
			"TestSomeThing3") {
			return
		}
	}
}

func TestOnlyValidTestMethodNamesAreExtracted(t *testing.T) {
	file := createFile(`
			package sometest
			import "testing"
			
			func Test(t *testing.T){}
			func TestActualTest(t *testing.T){}
			func Test_ThisToo(_ *testing.T){}
			
			func TestnoTest(t *testing.T){}
			func TestNotThis(t *T){}
			func TestNotTypeT(t *testing.U){}
			func TestNotAPointer(t testing.T){}
			func testNotCapitalTest(t testing.T){}
			func TesTestIsMissingThe_t(t testing.T){}
			func TestFewManyArguments(){}
			func TestTooManyArguments(t *testing.T, t *testing.T){}
			func TestMustNotHaveReturnType(t *testing.T) int {}
			var TestNoFunction int`)
	defer deleteFile(file)

	checkExtractsNames(t, file.Name(),
		"Test",
		"TestActualTest",
		"Test_ThisToo")
}

func TestOnErrorDuringExtraction_ResultIsEmptyArrayAndError(t *testing.T) {
	names, err := NewTestNameExtractor().ExtractFromFile("invalid_file_name.:///")
	if err == nil {
		t.Error("error was nil")
	}
	checkNames(t, names)
}

func createFile(content string) (file *os.File) {
	f, err := ioutil.TempFile("", "test2doc_test_")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(f.Name(), []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
	return f
}

func deleteFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}

func checkNames(t *testing.T, names []string, expected ...string) bool {
	if len(names) != len(expected) {
		t.Error(len(expected), "names expected, but they were", names)
		return false
	}
	for i := range expected {
		if names[i] != expected[i] {
			t.Error(expected, "expected but names were", names)
			return false
		}
	}
	return true
}

func checkExtractsNames(t *testing.T, fileName string, expected ...string) (ok bool) {
	names, err := NewTestNameExtractor().ExtractFromFile(fileName)
	if err != nil {
		t.Error("error was not nil")
	}
	return checkNames(t, names, expected...)
}
