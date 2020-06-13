// Package bob implements a solution to the exercism.io prompt by the same name
package bob

/* Hey returns the proper response for Bob's "personality" as defined in the prompt
Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.
*/
func Hey(remark string) string {
	//if remark to uppercase == remark then the person is yelling
	//if punctuation (last character) is a period
	//answer with 'Whoa, chill out!'
	//if punctuation is a question mark
	//answer with 'Calm down, I know what I'm doing!'
	//if remark is not yelling (else)
	//if punctuation is a question mark
	//answer with 'Sure.'
	//if remark is blank
	//answer with 'Fine. Be that way!'

	//otherwise answer with 'Whatever'
	return ""
}
