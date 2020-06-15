//Package collatzconjecture implements the exercism.io prompt by the same name
package collatzconjecture

import "fmt"

/*CollatzConjecture implements the logic of the mathematics principle by the same name.
which states that from any positive integer starting number
	1. if you divide an even number by 2
	2. and multiply an odd number by 3 and add 1
	3. and continue to iterate,
	you will always end up at 1 for a result, varying only in the steps required to reach that point
*/
func CollatzConjecture(start int) (steps int, e error) {
	if start <= 0 {
		return -1, fmt.Errorf("invalid starting number: %d", start)
	}

	for {
		if start == 1 {
			return
		}
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}

		steps++
	}
}
