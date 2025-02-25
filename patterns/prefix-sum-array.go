package patterns

func prefixSumArray(nums []int) []int {
	if nums == nil {
		return nil
	}

	if len(nums) == 0 {
		return []int{}
	}

	prefixSumArr := make([]int, len(nums))
	prefixSumArr[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSumArr[i] = nums[i] + prefixSumArr[i-1]
	}

	return prefixSumArr
}
