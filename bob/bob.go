// Package bob implements a solution to the exercism.io prompt by the same name
package bob

import (
	"regexp"
	"strings"
)

/* Hey returns the proper response for Bob's "personality" as defined in the prompt
Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.
*/
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion, _ := regexp.MatchString("\\A.*\\?\\z", remark)
	isYelling, _ := regexp.MatchString("\\A[A-Z1-9\\s[:punct:]]+[A-Z][A-Z1-9\\s[:punct:]]+\\z", remark)
	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYelling:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
