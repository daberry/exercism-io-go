//Package listops implements the suite of functions to operate on lists from exercism.io prompt by the same  name
package listops

import "container/list"

//IntList is a list of integers, implemented so that the test suite can function
type IntList list.List

//Length returns the number of elements contained within the IntList
func (il IntList) Length(listLength int) {
	_, ok := il.Front()
	for ok {

	}
	return
}
