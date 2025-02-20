/*
Given two strings `s` and `t`, return `true` if `t` is anagram of `s`, and `false` otherwise. An `anagram` is a word or phrase formed by
rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true
Example 2:
	Input: s = "rat", t = "car"
	Output: false
*/

package arrays_and_hashing

import (
	"sort"
)

func IsAnagram(s, t string) bool {
	if (len(s) != len(t)) || (len(s) < 1 || len(s) > 50000) || (len(t) < 1 || len(t) > 50000) {
		return false
	}

	m := make([]rune, 256)

	for _, letter := range s {
		m[letter] += 1
	}

	for _, letter := range t {
		m[letter] -= 1
		if m[letter] < 0 {
			return false
		}
	}

	return true
}

func IsAnagram_2(s, t string) bool {
	if (len(s) != len(t)) || (len(s) < 1 || len(s) > 50000) || (len(t) < 1 || len(t) > 50000) {
		return false
	}

	sbyte, tbyte := []byte(s), []byte(t)

	sort.Slice(sbyte, func(i, j int) bool {
		return sbyte[i] < tbyte[j]
	})

	sort.SliceStable(tbyte, func(i, j int) bool {
		return tbyte[i] < tbyte[j]
	})

	for idx, char := range sbyte {
		if char != tbyte[idx] {
			return false
		}
	}

	return true
}

func IsAnagram_3(s, t string) bool {
	if (len(s) != len(t)) || (len(s) < 1 || len(s) > 50000) || (len(t) < 1 || len(t) > 50000) {
		return false
	}

	m := make(map[rune]int, len(s)+len(t))

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		m[char]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
