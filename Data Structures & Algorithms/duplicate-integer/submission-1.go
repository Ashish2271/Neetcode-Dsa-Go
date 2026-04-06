func hasDuplicate(nums []int) bool {

	freq := make(map[int]int)
	for _,v := range nums {
		freq[v]++
	}
	for _ ,v := range freq{
		if v > 1{
			return true
		}
	
	}
	return false
}
