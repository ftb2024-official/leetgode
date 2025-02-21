/*
Given a string `s` consisting of lowercase English letters, return the first letter to appear twice. A letter `a` appears twice before another
letter`b` if the second occurence of `a` is before the second occurence of `b`. `s` will contain at least one letter that appears twice.

Example 1:
	Input: s = "abccbaacz"
	Output: "c"
	Explanation:
		The letter 'a' appears on the indexes 0, 5 and 6.
		The letter 'b' appears on the indexes 1 and 4.
		The letter 'c' appears on the indexes 2, 3 and 7.
		The letter 'z' appears on the index 8.
		The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.

Example 2:
	Input: s = "abcdd"
	Output: "d"
	Explanation: The only letter that appears twice is 'd' so we return 'd'.
*/

package arrays_and_hashing

func RepeatedCharacter(s string) byte {
	seen := map[rune]struct{}{}

	for _, char := range s {
		if _, ok := seen[char]; ok {
			return byte(char)
		}

		seen[char] = struct{}{}
	}

	return byte(0)
}

func RepeatedCharacter_2(s string) rune {
	seen := map[rune]bool{}

	for _, char := range s {
		if seen[char] {
			return char
		}

		seen[char] = true
	}

	return rune(0)
}
