package test2doc

import "testing"

func TestTheWordTestIsStrippedFromTheFront(t *testing.T) {
	words := c.Chop("TestFunction")
	checkWords(t, words, "Function")
}

func TestNamesAreSplitAtCapitalLetters(t *testing.T) {
	words := c.Chop("TestSomeFunction")
	checkWords(t, words, "Some", "Function")
}

func TestConsecutiveCapitalLettersAreEachTreatedAsOneWord(t *testing.T) {
	words := c.Chop("TestXYCoordinateBecomesXYZ")
	checkWords(t, words, "X", "Y", "Coordinate", "Becomes", "X", "Y", "Z")
}

func TestWordsAreSplitOnUnderscoreWhichIsThenDiscarded(t *testing.T) {
	words := c.Chop("Test_Some_other___Function")
	checkWords(t, words, "Some", "other", "Function")
}

func TestDigitsAreKeptTogetherAsOneWord(t *testing.T) {
	words := c.Chop("Test123Plus5equals128")
	checkWords(t, words, "123", "Plus", "5", "equals", "128")
}

func TestUnicodeLikeÄéßIsSupported(t *testing.T) {
	words := c.Chop("TestÄöüßÉÂâ")
	checkWords(t, words, "Äöüß", "É", "Ââ")
}

var c = CamelCaseChopper{}

func checkWords(t *testing.T, words []string, expected ...string) {
	if len(words) != len(expected) {
		t.Error(expected, "expected but words are", words)
	} else {
		for i := range expected {
			if expected[i] != words[i] {
				t.Error("word", i, "is wrong, words are", words)
			}
		}
	}
}
