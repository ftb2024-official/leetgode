package patterns

// Maximum Sum of a Subarray of Size `k`
func slidingWindow_max_sum_of_subarray_of_size_k(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return -1
	}

	var windowSum int

	for i := range k {
		windowSum += nums[i]
	}
	maxSum := windowSum

	for i := k; i < n; i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func slidingWindow_longest_substr_with_k_distinct_chars(str string, k int) int {
	return -1
}
