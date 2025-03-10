/*
Given an array of strings `words` and an integer `k`, return the `k` most frequent strings. Return the answer sorted by the frequency
from highest to lowest. Sort the words with the same frequency by theirlexicographic order.

Example 1:
	Input: words = ["i","love","leetcode","i","love","coding"], k = 2
	Output: ["i","love"]
	Explanation: "i" and "love" are the two most frequent words.
	Note that "i" comes before "love" due to a lower alphabetical order.

Example 2:
	Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
	Output: ["the","is","sunny","day"]
	Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.
*/

package yandex

import "sort"

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	var uniq []string

	for i := range words {
		if _, exists := freq[words[i]]; !exists {
			uniq = append(uniq, words[i])
		}

		freq[words[i]]++
	}

	sort.Slice(words, func(i int, j int) bool {
		if freq[words[i]] == freq[words[j]] {
			return uniq[i] < uniq[j]
		}

		return freq[uniq[i]] > freq[uniq[j]]
	})

	return uniq
}
