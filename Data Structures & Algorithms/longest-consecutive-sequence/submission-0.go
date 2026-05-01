func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxLen := 1
	currLen := 1

	for i := 1; i < len(nums); i++ {

		// skip duplicates
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == nums[i-1]+1 {
			currLen++
		} else {
			maxLen = Max(maxLen, currLen)
			currLen = 1
		}
	}

	return Max(maxLen, currLen)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}