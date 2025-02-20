/*
Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

Example 1:
	Input: nums = [1,2,3,1]
	Output: true
Example 2:
	Input: nums = [1,2,3,4]
	Output: false
Example 3:
	Input: nums = [1,1,1,3,3,4,3,2,4,2]
	Output: true
*/

package arrays_and_hashing

func ContainsDuplicate(nums []int) bool {
	dist := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := dist[num]; ok {
			return true
		}

		dist[num] = struct{}{}
	}

	return false
}
