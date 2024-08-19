//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	var runes []rune

	for _, r := range input {
		if r == utf8.RuneError {
			runes = append(runes, '\uFFFD')
		} else {
			runes = append(runes, r)
		}
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	var builder strings.Builder
	builder.Grow(len(input))

	for _, r := range runes {
		builder.WriteRune(r)
	}

	return builder.String()
}
