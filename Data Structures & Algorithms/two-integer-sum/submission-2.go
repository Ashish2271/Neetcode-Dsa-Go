func twoSum(nums []int, target int) []int {
	
    arrm := make(map[int]int)
	for i,v := range nums {
	
		arrm[v] = i 
     
	}
for i,v := range nums {
	ans := target - v
	if val,ok := arrm[ans]; ok && i!= val{
		
		return []int{i, val}
	}
     
	}
	return []int{}
}
