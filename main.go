package main

import (
	"fmt"
	test "leetgode/arrays_and_hashing"
)

func main() {
	fmt.Println(test.RemoveAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
	fmt.Println(test.RemoveAnagrams([]string{"a", "b", "c", "d", "e"}))
}
