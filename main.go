package main

import (
	"fmt"
	test "leetgode/arrays_and_hashing"
)

func main() {
	fmt.Println(test.FindKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
	fmt.Println(test.FindKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
}
