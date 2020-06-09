// Package scrabble implements a solution to the exercism.io problem by the same name
package scrabble

import "strings"

/*Score returns the scrabble score for the given string, by assigning the following values to letters:
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10

*/
func Score(input string) (score int) {
	//convert string to lowercase
	input = strings.ToLower(input)
	//iterate through characters in string
	for _, val := range input {
		switch val {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	//return result
	return score
}
