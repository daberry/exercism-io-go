/*
Package raindrops implements a solution the exercism.io problem by the same name
*/
package raindrops

import "strconv"

/*
Convert returns a string that follows this logic for the input integer:
if 3 is a factor of the input, concatenate "Pling" to the result
if 5 is a factor of the input, concatenate "Plang" to the result
if 7 is a factor of the input, concatenate "Plong" to the result
if none of the above are factors of the input, return the input number as a string
*/
func Convert(num int) (result string) {
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(num)
	}

	return result
}
