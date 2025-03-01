/*
This technique should be your go-to when you see a question that involves searching for a pair (or more) of elements in an array that meet a
certain criteria.
*/

package patterns

import "sort"

func twoPointer_sorted_array(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	if len(nums) < 2 {
		return []int{}
	}

	sort.Ints(nums)

	left, right := 0, len(nums)-1
	var currentSum int

	for left != right {
		currentSum = nums[left] + nums[right]
		if currentSum == target {
			return []int{nums[left], nums[right]}
		} else if currentSum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}

func twoPointer_palindrome(str string) bool {
	if len(str) == 0 {
		return false
	}

	if len(str) == 1 {
		return true
	}

	left, right := 0, len(str)-1

	for left <= right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}
