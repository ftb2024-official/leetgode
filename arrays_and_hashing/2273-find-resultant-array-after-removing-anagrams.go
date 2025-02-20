/*
You are given a 0-indexed string array `words`, where `words[i]` consists of lowercase English letters.

In one operation, select any index `i` such that `0 < i < words.length` and `words[i - 1]` and `words[i]` are anagrams, and delete `words[i]` from
`words`. Keep performing this operation as long as you can select an index that satisfies the conditions.

Return `words` after performing all operations. It can be shown that selecting the indices for each operation in any arbitrary order will lead to
the same result.

Example 1:
	Input: words = ["abba","baba","bbaa","cd","cd"]
	Output: ["abba","cd"]
	Explanation:
		One of the ways we can obtain the resultant array is by using the following operations:
		- Since words[2] = "bbaa" and words[1] = "baba" are anagrams, we choose index 2 and delete words[2].
		Now words = ["abba","baba","cd","cd"].
		- Since words[1] = "baba" and words[0] = "abba" are anagrams, we choose index 1 and delete words[1].
		Now words = ["abba","cd","cd"].
		- Since words[2] = "cd" and words[1] = "cd" are anagrams, we choose index 2 and delete words[2].
		Now words = ["abba","cd"].
		We can no longer perform any operations, so ["abba","cd"] is the final answer.

Example 2:
	Input: words = ["a","b","c","d","e"]
	Output: ["a","b","c","d","e"]
	Explanation:
		No two adjacent strings in words are anagrams of each other, so no operations are performed.
*/

package arrays_and_hashing

import (
	"sort"
	"strings"
)

func sortStr(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}

// func sortStr(str string) string {
// 	bytedStr := []byte(str)
// 	sort.SliceStable(bytedStr, func(a, b int) bool { return bytedStr[a] < bytedStr[b] })
// 	return string(bytedStr)
// }

func RemoveAnagrams(words []string) []string {
	for i := len(words) - 1; i > 0; i-- {
		j := i - 1

		if sortStr(words[i]) == sortStr(words[j]) {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}
