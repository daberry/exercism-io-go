// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb implements a solution to the exercism.io prompt by the same name
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb returns an array of assembled strings
func Proverb(rhyme []string) (assembledProverb []string) {
	if len(rhyme) == 0 {
		return
	}

	assembledProverb = make([]string, len(rhyme))
	for i, val := range rhyme {
		if i+1 < len(rhyme) {
			assembledProverb[i] = fmt.Sprintf(stanza, val, rhyme[i+1])
		} else {
			assembledProverb[i] = fmt.Sprintf(last, rhyme[0])
		}
	}
	return
}
