package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var prev rune
	var escape bool
	var result strings.Builder
	for _, s := range str {
		if escape {
			addRune(1, prev, &result)
			if !(unicode.IsDigit(s) || s == '\\') {
				return "", ErrInvalidString
			}
			prev = s
			escape = false
			continue
		}
		if s == '\\' {
			escape = true
			continue
		}
		if unicode.IsDigit(s) {
			if prev == 0 {
				return "", ErrInvalidString
			}
			count := int(s - '0')
			addRune(count, prev, &result)
			prev = 0
			continue
		}

		addRune(1, prev, &result)
		prev = s
	}
	addRune(1, prev, &result)
	return result.String(), nil
}

func addRune(count int, char rune, result *strings.Builder) {
	if char == 0 {
		return
	}
	for i := 0; i < count; i++ {
		result.WriteRune(char)
	}
}
