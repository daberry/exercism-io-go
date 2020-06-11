//Package etl implements a solution to the exercism.io prompt by the same name
package etl

import "strings"

//Transform modifies the scrabble scoring format to fall in line with the new requirements
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)

	for i, letters := range input {
		for _, letter := range letters {
			letter = strings.ToLower(letter)
			output[letter] = i
		}
	}

	return
}
