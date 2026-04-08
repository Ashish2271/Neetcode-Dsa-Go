func twoSum(nums []int, target int) []int {

	arrm := make(map[int]int)

	for i, v := range nums {

		ans := target - v

		// check first
		if j, ok := arrm[ans]; ok {
			return []int{j, i}
		}

		// store after check
		arrm[v] = i
	}

	return []int{}
}