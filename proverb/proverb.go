// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb implements a solution to the exercism.io prompt by the same name
package proverb

// Proverb returns an array of assembled strings
func Proverb(rhyme []string) (assembled []string) {
	for i, val := range rhyme {
		if i < len(rhyme)-1 {
			assembled = append(assembled, "For want of a "+val+" the "+rhyme[i+1]+" was lost.")
		} else {
			assembled = append(assembled, "And all for the want of a "+rhyme[0]+".")
		}
	}
	return
}
