//Package romannumerals implements a solution to the exercism.io prompt by the same name
package romannumerals

import (
	"fmt"
	"strings"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

var arabicToRomanDictionary []arabicToRoman = []arabicToRoman{
	{1000, "M"},
	{900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
	{90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"},
	{9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

//ToRomanNumeral converts an Arabic numeral integer to a Roman numeral string
func ToRomanNumeral(number int) (string, error) {
	var str strings.Builder

	if number <= 0 || number > 3000 {
		return "", fmt.Errorf("provided arabic integer (%d) is outside of acceptable range", number)
	}

	for _, entry := range arabicToRomanDictionary {
		for number >= entry.arabic {
			fmt.Println(entry.arabic, number)
			str.WriteString(entry.roman)
			number -= entry.arabic
		}
	}

	return str.String(), nil
}
