/*
Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except
`nums[i]`. The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

Example 1:
	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]

Example 2:
	Input: nums = [-1,1,0,-3,3]
	Output: [0,0,9,0,0]
*/

package neetcode

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, len(nums))

	left := 1
	for i := 0; i < n; i++ {
		answer[i] = left
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}
