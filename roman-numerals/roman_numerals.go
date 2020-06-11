//Package romannumerals implements a solution to the exercism.io prompt by the same name
package romannumerals

import "errors"

//ToRomanNumeral converts an Arabic numeral integer to a Roman numeral string
func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("provided arabic integer outside of acceptable range")
	}
	for ok := true; ok; ok = arabic > 0 {
		if arabic >= 1000 {
			roman += "M"
			arabic -= 1000
		} else if arabic >= 900 {
			roman += "CM"
			arabic -= 900
		} else if arabic >= 500 {
			roman += "D"
			arabic -= 500
		} else if arabic >= 400 {
			roman += "CD"
			arabic -= 400
		} else if arabic >= 100 {
			roman += "C"
			arabic -= 100
		} else if arabic >= 90 {
			roman += "XC"
			arabic -= 90
		} else if arabic >= 50 {
			roman += "L"
			arabic -= 50
		} else if arabic >= 40 {
			roman += "XL"
			arabic -= 40
		} else if arabic >= 10 {
			roman += "X"
			arabic -= 10
		} else if arabic >= 9 {
			roman += "IX"
			arabic -= 9
		} else if arabic >= 5 {
			roman += "V"
			arabic -= 5
		} else if arabic >= 4 {
			roman += "IV"
			arabic -= 4
		} else {
			roman += "I"
			arabic--
		}
	}
	return roman, err
}
