/*
Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one. You must implement a solution with
a linear runtime complexity and use only constant extra space.

Example 1:
	Input: nums = [2,2,1]
	Output: 1

Example 2:
	Input: nums = [4,1,2,1,2]
	Output: 4

Example 3:
	Input: nums = [1]
	Output: 1
*/

package yandex

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}

// XOR approach
func singleNumber_2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var result int

	for _, num := range nums {
		result ^= num
	}

	return result
}
