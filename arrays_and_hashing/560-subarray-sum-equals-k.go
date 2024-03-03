/*
Given an array of integers `nums` and an integer `k`, return the total number of subarrays whose sum equals to `k`. A subarray is a contiguous
non-empty sequence of elements within an array.

Example 1:
	Input: nums = [1,1,1], k = 2
	Output: 2

Example 2:
	Input: nums = [1,2,3], k = 3
	Output: 2
*/

package arraysandhashing

func SubarraySum(nums []int, k int) int {
	answer, sum, prefixSum := 0, 0, map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefixSum[sum-k]; ok {
			answer += prefixSum[sum-k]
		}

		prefixSum[sum] += 1
	}

	return answer
}
