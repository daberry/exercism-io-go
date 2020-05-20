package scale

var (
	sharpPitches         = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flatPitches          = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	noSharpOrFlatPitches = []string{"C", "D", "E", "F", "G", "A", "B"}

	sharpTonics         = []string{"G", "D", "A", "E", "B", "F# major", "e", "b", "f#", "c#", "g#", "d# minor"}
	flatTonics          = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb major", "d", "g", "c", "f", "bb", "eb minor"}
	noSharpOrFlatTonics = []string{"C major", "a minor"}
)

type ()

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
	return sharpPitches
}
