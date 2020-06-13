//Package dna implements a solution to the exercism.io prompt 'nucleotide_count'
package dna

import (
	"fmt"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	if isInvalid, _ := regexp.MatchString(`[^ATGC]`, string(d)); isInvalid {
		return nil, fmt.Errorf("invalid character in nucleotide string")
	}

	var h Histogram = map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for nucleotide := range h {
		for _, rune := range d {
			if nucleotide == rune {
				h[nucleotide]++
			}
		}
	}

	return h, nil
}
