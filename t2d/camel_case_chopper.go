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
// reads a single rune and outputs the next state or nil to stay in the same
// state and keep reading.
type readFunc func(r rune) (changeStateTo readFunc)

// The string is iterated rune by rune. Each rune is passed to the current state
// readFunc. If it changes state (return non-nil value) the word is chopped at
// the current position.
func (reader nameReader) read(name string) []string {
	words := make([]string, 0, 20)
	start := 0
	for i, r := range name {
		if next := reader.handler(r); next != nil {
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

func readWord(r rune) readFunc {
	if r == '_' {
		return skipUnderscore
	}
	if unicode.IsDigit(r) {
		return readNumber
	}
	if unicode.IsUpper(r) {
		return readWord
	}
	return nil
}

func readNumber(r rune) readFunc {
	if r == '_' {
		return skipUnderscore
	}
	if unicode.IsDigit(r) {
		return nil
	}
	return readWord
}

func skipUnderscore(r rune) readFunc {
	if r == '_' {
		return skipUnderscore
	}
	if unicode.IsDigit(r) {
		return readNumber
	}
	return readWord
}

func isAllUnderscores(word string) bool {
	return word == strings.Repeat("_", len(word))
}
