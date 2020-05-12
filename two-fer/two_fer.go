// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith should have a comment documenting it.
/*
Given a name, return a string with the message:
One for X, one for me.

Where X is the given name.

However, if the name is missing, return the string:
One for you, one for me.


Here are some examples:

|Name    |String to return
|:-------|:------------------
|Alice   |One for Alice, one for me.
|Bob     |One for Bob, one for me.
|        |One for you, one for me.
|Zaphod  |One for Zaphod, one for me.

*/
func ShareWith(name string) string {
	if (name == "") {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."
}
