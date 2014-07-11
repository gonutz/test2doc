package t2d

import (
	"strings"
	"unicode"
)

type CamelCaseChopper struct{}

func NewCamelCaseChopper() CamelCaseChopper { return CamelCaseChopper{} }

func (c CamelCaseChopper) Chop(name string) []string {
	nameOnly := strings.TrimPrefix(name, "Test")
	finalWordDelimiter := "A"
	return nameReader{readWord}.read(nameOnly + finalWordDelimiter)
}

type nameReader struct {
	handler readFunc
}

// The nameReader is really a state machine and the readFunc is a state. It
// reads a single rune and outputs the next state and indicates whether the word
// needs to be chopped at the current position
type readFunc func(r rune) (nextState readFunc, chopWord bool)

func (reader nameReader) read(name string) []string {
	words := make([]string, 0, 20)
	start := 0
	for i, r := range name {
		if next, chopWord := reader.handler(r); chopWord {
			reader.handler = next
			word := name[start:i]
			start = i
			if !isAllUnderscores(word) {
				words = append(words, word)
			}
		}
	}
	return words
}

func readWord(r rune) (nextState readFunc, chopWord bool) {
	if r == '_' {
		return skipUnderscore, true
	}
	if unicode.IsDigit(r) {
		return readNumber, true
	}
	if unicode.IsUpper(r) {
		return readWord, true
	}
	return readWord, false
}

func readNumber(r rune) (nextState readFunc, chopWord bool) {
	if r == '_' {
		return skipUnderscore, true
	}
	if unicode.IsDigit(r) {
		return readNumber, false
	}
	return readWord, true
}

func skipUnderscore(r rune) (nextState readFunc, chopWord bool) {
	if r == '_' {
		return skipUnderscore, true
	}
	if unicode.IsDigit(r) {
		return readNumber, true
	}
	return readWord, true
}

func isAllUnderscores(word string) bool {
	return word == strings.Repeat("_", len(word))
}
