package t2d

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
	fileName := createFile(`
			package sometest
			import "testing"

			func TestSomeThing1(t *testing.T){}
			func TestSomeThing2(t *testing.T){}
			func TestSomeThing3(t *testing.T){}`)
	defer deleteFile(fileName)

	for i := 0; i < 100; i++ {
		if !checkExtractsNames(t, fileName,
			"TestSomeThing1",
			"TestSomeThing2",
			"TestSomeThing3") {
			return
		}
	}
}

func TestOnlyValidTestMethodNamesAreExtracted(t *testing.T) {
	fileName := createFile(`
			package sometest
			import "testing"
			
			func Test(t *testing.T){}
			func TestActualTest(t *testing.T){}
			func Test_ThisToo(_ *testing.T){}
			
			func TestfunctionNameDoesNotHaveCapitalLetterAfterTest(t *testing.T){}
			func TestTNotInTestingPackage(t *T){}
			func TestNotTypeT(t *testing.U){}
			func TestNotAPointer(t testing.T){}
			func testNotCapitalTest(t *testing.T){}
			func TesTestIsMissingThe_t_inTheWordTest(t *testing.T){}
			func TestMissingArgument(){}
			func TestTooManyArguments(t, u *testing.T){}
			func TestMustNotHaveReturnType(t *testing.T) int {}
			var TestNoFunction int`)
	defer deleteFile(fileName)

	checkExtractsNames(t, fileName,
		"Test",
		"TestActualTest",
		"Test_ThisToo")
}

func TestOnErrorDuringExtraction_ResultIsEmptyArrayAndError(t *testing.T) {
	names, err := NewTestNameExtractor().ExtractTestsFromFile("invalid_file_name.:///")
	if err == nil {
		t.Error("error was nil")
	}
	checkNames(t, names)
}

func createFile(content string) string {
	f, err := ioutil.TempFile("", "test2doc_test_")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(f.Name(), []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return f.Name()
}

func deleteFile(path string) {
	os.Remove(path)
}

func checkNames(t *testing.T, names []string, expected ...string) bool {
	if len(names) != len(expected) {
		t.Error(len(expected), "name(s) expected, but they were", names)
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
	names, err := NewTestNameExtractor().ExtractTestsFromFile(fileName)
	if err != nil {
		t.Error("error was not nil")
	}
	return checkNames(t, names, expected...)
}
