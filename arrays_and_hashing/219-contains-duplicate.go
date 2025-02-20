/*
Given an integer array `nums` and an integer `k`, return `true` if there are two distinct indicies `i` and `j` in the array such that
`num[i] = num[j]` and `abs(i - j) <= k`.

Example 1:
	Input: nums = [1,2,3,1], k = 3
	Output: true
Example 2:
	Input: nums = [1,0,1,1], k = 1
	Output: true
Example 3:
	Input: nums = [1,2,3,1,2,3], k = 2
	Output: false
*/

package arrays_and_hashing

func ContainsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)

	for idx, num := range nums {
		if val, ok := hash[num]; ok {
			if idx-val <= k {
				return true
			}
		}

		hash[num] = idx
	}

	return false
}
