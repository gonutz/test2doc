package t2d

import (
	"strings"
	"unicode"
)

type CamelCaseChopper struct{}

func (c CamelCaseChopper) Chop(name string) []string {
	nameOnly := strings.TrimPrefix(name, "Test")
	finalWordDelimiter := "A"
	return nameReader{readWord}.read(nameOnly + finalWordDelimiter)
}

type readFunc func(r rune) (changeStateTo readFunc)

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

type nameReader struct {
	handler readFunc
}

func isAllUnderscores(word string) bool {
	return word == strings.Repeat("_", len(word))
}

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
