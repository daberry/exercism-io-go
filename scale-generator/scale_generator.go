package scale

import (
	"fmt"
)

var (
	sharpPitches         = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flatPitches          = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	noSharpOrFlatPitches = []string{"C", "D", "E", "F", "G", "A", "B"}

	sharpTonics         = []string{"G", "D", "A", "E", "B", "F# major", "e", "b", "f#", "c#", "g#", "d# minor"}
	flatTonics          = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb major", "d", "g", "c", "f", "bb", "eb minor"}
	noSharpOrFlatTonics = []string{"C major", "a minor"}
)

type ()

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func getNextTone(currentTone string, intervalToNextTone string, tonic string) string {
	var distanceToNextTone int

	if intervalToNextTone == "M" {
		distanceToNextTone = 2
	}

	if intervalToNextTone == "m" {
		distanceToNextTone = 1
	}

	if intervalToNextTone == "" {
		distanceToNextTone = -1
	}

	return "placeholder" + string(distanceToNextTone)
}

//Scale returns the appropriate scale for the given tonic and interval
func Scale(tonic string, interval string) []string {
	resultScale := make([]string, len(sharpPitches))
	fmt.Println(resultScale)
	if tonic == "C" {
		return sharpPitches
	}

	fmt.Println(reflect.typeof(resultScale[5]))

	if Contains(flatTonics, tonic) {
		copy(resultScale, flatPitches)
		start := Find(resultScale, tonic)
		fmt.Println(reflect.typeof(resultScale[start]))
		return reflect.typeof(resultScale[start:])
		// resultScale = append(resultScale[start:], resultScale[:start])
	}

	return resultScale
}
