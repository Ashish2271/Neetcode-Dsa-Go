func majorityElement(nums []int) int {
    major := make(map[int]int)
	for _,v := range nums{
		major[v]++
	}
    ans := 0
	maxval :=0
	for k,v := range major {
     if v > maxval {
		maxval = v
		ans = k
	 }
	}
 return ans 
}
