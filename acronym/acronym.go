// Package acronym implements a solution to the exercism.io prompt by the same name
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns an all capital letter acronym for the input string, delimiting words with non-word characters as defined by Regex
func Abbreviate(s string) (acronym string) {
	stripDashes := regexp.MustCompile(`-`).ReplaceAllString(s, " ")
	noPunctuation := regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(stripDashes, "")
	trimmed := regexp.MustCompile(`\s{2,}`).ReplaceAllString(noPunctuation, " ")

	words := strings.Split(trimmed, " ")

	for _, word := range words {
		if len(word) > 0 {
			acronym += string(word[0])
		}
	}

	return strings.ToUpper(acronym)
}
