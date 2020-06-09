// Package leap implements a solution the exercism.io problem Leap Year
package leap

/*IsLeapYear returns a boolean determining if the given integer year is a leap year via the following rules:
on every year that is evenly divisible by 4
	except every year that is evenly divisible by 100
		unless the year is also evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
