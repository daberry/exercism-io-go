// Then complete the exercise by implementing these methods:
//
//    (Ints) Keep(func(int) bool) Ints
//    (Ints) Discard(func(int) bool) Ints
//    (Lists) Keep(func([]int) bool) Lists
//    (Strings) Keep(func(string) bool) Strings
package strain

import (
	"reflect"
)

//Ints 1d integer list
type Ints []int

//Lists 2d integer list
type Lists [][]int

//Strings 1d string list
type Strings []string

//Keep filters the given list and returns the items which pass the predicate test
func (list Ints) Keep(predicate func(int) bool) (keepers Ints) {
	for _, item := range list {
		if predicate(item) {
			keepers = append(keepers, item)
		}
	}

	return
}

//Discard filters the given list and returns the items which fail the predicate test
func (list Ints) Discard(predicate func(int) bool) (discards Ints) {
	if len(list) == 0 {
		return
	}

	for _, item := range list {
		if !predicate(item) {
			discards = append(discards, item)
		}
	}

	return
}

//Keep filters the given list and returns the items which pass the predicate test
func (lists Lists) Keep(predicate func([]int) bool) (keepers Lists) {
	for _, list := range lists {
		if !reflect.ValueOf(list).IsNil() && predicate(list) {
			keepers = append(keepers, list)
		}
	}

	return
}

//Keep filters the given list and returns the items which pass the predicate test
func (strings Strings) Keep(predicate func(string) bool) (keepers Strings) {
	for _, item := range strings {
		if predicate(item) {
			keepers = append(keepers, item)
		}
	}

	return
}
