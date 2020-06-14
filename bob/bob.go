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

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	re, _ := regexp.Compile(`[a-zA-Z]`)

	punctuation := remark[len(remark)-1]

	if re.MatchString(remark) && strings.ToUpper(remark) == remark {
		if punctuation == '?' {
			return "Calm down, I know what I'm doing!"
		}

		return "Whoa, chill out!"
	}

	if punctuation == '?' {
		return "Sure."
	}

	return "Whatever."
}
