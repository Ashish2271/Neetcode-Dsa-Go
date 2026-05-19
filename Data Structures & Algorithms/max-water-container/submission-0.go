func maxArea(heights []int) int {
    maxArea := 0
    currArea := 0
    width := 0

 
    for i :=0; i<len(heights) - 1; i++{
        for j :=1; j<len(heights); j++{
            width = j-i
            currArea = min(heights[i],heights[j]) * width
            maxArea =  max(maxArea,currArea) 
        }
    } 
return maxArea
}
