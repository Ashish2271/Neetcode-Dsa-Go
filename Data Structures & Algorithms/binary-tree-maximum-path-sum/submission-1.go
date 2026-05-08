/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := -1 << 31
	dfs(root,&res)
	return res 

}

func dfs(root *TreeNode, res *int){
	if root == nil {
		return 
	}
	left := getMax(root.Left)
	right := getMax(root.Right)
	*res = max(*res,root.Val + left + right)
	dfs(root.Left,res)
	dfs(root.Right,res)
}


func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMax(root.Left)
	right := getMax(root.Right)
	path := root.Val + max (left,right)
	return max(0,path)
}