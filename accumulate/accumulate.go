//Package accumulate implements a solution to the exercism.io problem with the same name
package accumulate

//Converter is a type definition for the provided iterator to map input to output
type Converter func(string) string

/*
Accumulate applies the provided iterator to the provided collection and
returns the result of the transformation
*/
func Accumulate(strings []string, conv Converter) (result []string) {
	result = make([]string, len(strings))

	for i, val := range strings {
		result[i] = conv(val)
	}

	return
}
