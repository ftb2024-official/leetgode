package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 2, 4}
	prefixNums := make([]int, len(nums))
	prefixNums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixNums[i] = nums[i] + prefixNums[i-1]
	}

	fmt.Println("nums:", nums)
	fmt.Println("prefix nums:", prefixNums)
}
