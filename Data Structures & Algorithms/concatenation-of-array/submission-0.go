func getConcatenation(nums []int) []int {
    n := len(nums)
    ans := make([]int,2*n)
    for  i := 0 ;  i < 2*n; i++ {
        if i < n {
            ans [i] = nums[i]
        }
        if i > n-1 {
            ans[i] = nums [i-n]
        } 
    }
    return ans 
}
