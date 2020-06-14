// Package acronym implements a solution to the exercism.io prompt by the same name
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

// Abbreviate returns an all capital letter acronym for the input string, delimiting words with non-word characters as defined by Regex
func Abbreviate(s string) (acronym string) {
	//split s using non-word char delimiter
	noPunctuation := regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(s, "")
	trimmed := regexp.MustCompile(`\s{2,}`).ReplaceAllString(noPunctuation, " ")
	fmt.Println("trimmed", trimmed)
	words := strings.Split(trimmed, " ")
	fmt.Println("words", words, len(words))
	//create string from first letter in every word
	for _, word := range words {
		fmt.Println("current word", word)
		if len(word) > 0 {
			acronym += string(word[0])
		}
	}
	//capitalize letters for acronym

	return strings.ToUpper(acronym)
}
