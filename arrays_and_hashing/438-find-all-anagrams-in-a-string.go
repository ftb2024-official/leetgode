/*
Given two strings `s` and `p`, return an array of all the start indices of `p`s anagrams in `s`. You may return the answer in any order.

Example 1:
	Input: s = "cbaebabacd", p = "abc"
	Output: [0,6]
	Explanation:
		The substring with start index = 0 is "cba", which is an anagram of "abc".
		The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:
	Input: s = "abab", p = "ab"
	Output: [0,1,2]
	Explanation:
		The substring with start index = 0 is "ab", which is an anagram of "ab".
		The substring with start index = 1 is "ba", which is an anagram of "ab".
		The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

package arrays_and_hashing

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func FindAnagrams(s, p string) []int {
	ls, lp := len(s), len(p)
	ans := make([]int, 0)
	pattern := make([]int, 26)

	for _, letter := range p {
		pattern[letter-'a']++
	}

	try := make([]int, 26)
	for i := 0; i < lp; i++ {
		try[s[i]-'a']++
	}

	if equal(try, pattern) {
		ans = append(ans, 0)
	}

	for next := lp; next < ls; next++ {
		del := next - lp
		try[s[del]-'a']--
		try[s[next]-'a']++

		if equal(try, pattern) {
			ans = append(ans, del+1)
		}
	}

	return ans
}
