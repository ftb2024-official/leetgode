/*
Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:
	Input: strs = [""]
	Output: [[""]]
Example 3:
	Input: strs = ["a"]
	Output: [["a"]]
*/

package arrays_and_hashing

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	kb := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)

		if _, ok := kb[key]; !ok {
			kb[key] = []string{}
		}

		kb[key] = append(kb[key], str)
	}

	res := make([][]string, len(kb))

	i := 0

	for _, anagrams := range kb {
		res[i] = anagrams
		i++
	}

	return res
}

func GroupAnagrams_2(strs []string) [][]string {
	anagramGroups := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}

		anagramGroups[count] = append(anagramGroups[count], str)
	}

	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func sortString(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
