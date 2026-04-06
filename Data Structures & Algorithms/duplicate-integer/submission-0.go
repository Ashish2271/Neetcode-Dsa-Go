func hasDuplicate(nums []int) bool {

	var ans  bool 
    for i := 0; i< len(nums); i++ {
		for j := i+1; j< len(nums); j++{
			if nums[i] == nums[j]{
				return true  
			}
			

		}
	}
	return ans
}
