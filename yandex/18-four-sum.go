/*
Given an array `nums` of `n` integers, return an array of all the unique quadruplets `nums[a], nums[b], nums[c], nums[d]` such that:
	- 0 <= a, b, c, d < n;
	- a, b, c, and d are distinct;
	- nums[a] + nums[b] + nums[c] + nums[d] == target;
You may return the answer in any order.

Example 1:
	Input: nums = [1,0,-1,0,-2,2], target = 0
	Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:
	Input: nums = [2,2,2,2,2], target = 8
	Output: [[2,2,2,2]]
*/

package yandex

import (
	"sort"
)

func fourSum_brute_force(nums []int, target int) [][]int {
	n := len(nums)
	uniqueQuadruplets := make(map[[4]int]bool, n)
	result := make([][]int, n)

	sort.Ints(nums)

	for a := 0; a < n-3; a++ {
		for b := a + 1; b < n-2; b++ {
			for c := b + 1; c < n-1; c++ {
				for d := c + 1; d < n; d++ {
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						quadruplet := [4]int{nums[a], nums[b], nums[c], nums[d]}
						uniqueQuadruplets[quadruplet] = true
					}
				}
			}
		}
	}

	for quad := range uniqueQuadruplets {
		result = append(result, quad[:])
	}

	return result
}
