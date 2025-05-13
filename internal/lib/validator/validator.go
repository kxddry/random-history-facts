package validator

import (
    "errors"
    "unicode"
)

var (
    ErrNonCyrillic = errors.New("non-cyrillic characters found")
)

func Validate(fact string) error {
    for _, ch := range fact {
        if !(isCyrillic(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) || unicode.IsPunct(ch)) {
            return ErrNonCyrillic
        }
    }
    return nil
}

func isCyrillic(ch rune) bool {
    return ('А' <= ch && ch <= 'Я') || ('а' <= ch && ch <= 'я')
}
