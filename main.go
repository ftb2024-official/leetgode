package main

import (
	"fmt"
	test "leetgode/arrays_and_hashing"
)

func main() {
	fmt.Println(test.ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(test.ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(test.ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
