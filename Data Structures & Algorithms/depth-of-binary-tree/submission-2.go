/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    maxDep := 0
     
      if root == nil {
        return 0
    }
    
    maxDep ++ 
    maxleft := maxDepth(root.Left)
    maxright := maxDepth(root.Right)
    maxDep = 1+max(maxleft,maxright)
 return maxDep
    
}
