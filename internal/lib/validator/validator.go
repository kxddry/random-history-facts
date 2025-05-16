package validator

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrNonCyrillic = errors.New("non-cyrillic characters found")
)

func roman(ch rune) bool {
	str := "IVXLCDM"
	return strings.Contains(str, string(ch))
}

func Validate(fact string) error {
	for _, ch := range fact {
		if !(roman(ch) || isCyrillic(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) || unicode.IsPunct(ch)) {
			return ErrNonCyrillic
		}
	}
	return nil
}

func isCyrillic(ch rune) bool {
	return ('А' <= ch && ch <= 'Я') || ('а' <= ch && ch <= 'я') || (ch == 'ё' || ch == 'Ё')
}
