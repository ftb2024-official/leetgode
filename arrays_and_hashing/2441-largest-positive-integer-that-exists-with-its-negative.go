/*
Given an integer array `nums` that does not contain any zeros, find the largest positive integer `k` such that `-k` also exists in the array.
Return the positive integer `k`. If there is no such integer, return `-1`.

Example 1:
	Input: nums = [-1,2,-3,3]
	Output: 3
	Explanation: 3 is the only valid k we can find in the array.

Example 2:
	Input: nums = [-1,10,6,7,-7,1]
	Output: 7
	Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.

Example 3:
	Input: nums = [-10,8,6,7,-2,-3]
	Output: -1
	Explanation: There is no a single valid k, we return -1.
*/

package arraysandhashing

import (
	"slices"
)

func FindMaxK(nums []int) int {
	slices.Sort(nums)
	posMap := map[int]struct{}{}

	for _, num := range nums { // map of only positive integers
		if num < 0 {
			continue
		}

		posMap[num] = struct{}{}
	}

	for _, num := range nums {
		if num > 0 {
			break
		}
		if _, ok := posMap[-num]; ok {
			return -num
		}
	}

	return -1
}

func FindMaxK_2(nums []int) int {
	max, posMap, negMap := -1, map[int]struct{}{}, map[int]struct{}{}

	for _, num := range nums {
		if num < 0 {
			negMap[num] = struct{}{}
			if _, ok := posMap[-num]; ok {
				if -num > max {
					max = -num
				}
			}
		} else {
			posMap[num] = struct{}{}
			if _, ok := negMap[-num]; ok {
				if num > max {
					max = num
				}
			}
		}
	}

	return max
}
