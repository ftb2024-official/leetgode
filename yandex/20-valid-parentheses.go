/*
Given a string `s` containing just the characters `(`, `(`, `{`, `}`, `[`, `]`, determine if the input string is valid.
An input string is valid if:
	- open brackets must be closed by the same type of brackets;
	- open brackets must be closed in the correct order;
	- every close bracket has a corresponding open bracket of the same type;

Example 1:
	Input: s = '()'
	Output: true

Example 2:
	Input: s = '()[]{}'
	Output: true

Example 3:
	Input: s = '(]'
	Output: false

Example 4:
	Input: s = '([])'
	Output: true
*/

package yandex

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	closingBrackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != closingBrackets[ch] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
