//Package twofer implements a solution to the exercism.io prompt with the same name.
package twofer

import "fmt"

/*ShareWith returns the proper twoFer result based on the name string provided
1. Given a name, return a string with the message:
	 One for X, one for me.
2. If no name is provided, return a string with the message:
   One for you, one for me.
*/
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
