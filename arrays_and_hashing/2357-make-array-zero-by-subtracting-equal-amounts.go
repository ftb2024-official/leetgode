/*
You are given a non-negative integer array `nums`. In one operation you must: 1) Choose a positive integer `x` such that `x` is less than or equal
to the smallest non-zero element in nums. 2) Subtract `x` from every positive element in nums.

Return the minimum number of operations to make every element in `nums` equal to `0`.

Example 1:
	Input: nums = [1,5,0,3,5]
	Output: 3
	Explanation:
		In the first operation, choose x = 1. Now, nums = [0,4,0,2,4].
		In the second operation, choose x = 2. Now, nums = [0,2,0,0,2].
		In the third operation, choose x = 2. Now, nums = [0,0,0,0,0].

Example 2:
	Input: nums = [0]
	Output: 0
	Explanation: Each element in nums is already 0 so no operations are needed.
*/

package arrays_and_hashing

func MinimumOperations(nums []int) int {
	bucket := make(map[int]struct{})

	for _, num := range nums {
		if num == 0 {
			continue
		}
		bucket[num] = struct{}{}
	}

	return len(bucket)
}
