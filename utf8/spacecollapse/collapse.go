//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var builder strings.Builder
	var lastWasSpace bool

	for _, r := range input {
		if r == utf8.RuneError {
			builder.WriteRune('\uFFFD')
			lastWasSpace = false
			continue
		}

		if unicode.IsSpace(r) {
			if !lastWasSpace {
				builder.WriteRune(' ')
				lastWasSpace = true
			}
		} else {
			builder.WriteRune(r)
			lastWasSpace = false
		}
	}

	return builder.String()
}
