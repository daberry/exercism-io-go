/*
Package hamming implements a solution to the exercism.io problem by the same name
*/
package hamming

import (
	"errors"
)

/*
Distance returns the Hamming distance between the two strings, which is defined as the number of characters that are different between the strings.
For example, consider the two following DNA strings:

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^

The DNA strand strings have 7 differences (denoted with ^), and therefore the Hamming distance is 7.

The Hamming distance is only defined for sequences of equal length. In this case 0 and an error are returned.
*/
func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return 0, errors.New("provided strings are not equal length")
	}

	if a != b {
		for i := range a {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
	}

	return hammingDistance, nil
}
