/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {

	arr := [][]int{}
    var dfs func(node *TreeNode,depth int) 
	dfs = func(node *TreeNode,depth int){
		if node == nil {
			return 
		}
        
       if len(arr) == depth {
            arr = append(arr, []int{})
        }

      arr[depth] = append(arr[depth],node.Val)

	  dfs(node.Left,depth+1)
	  dfs(node.Right,depth+1)

	}
	dfs(root,0)
	return arr 
}
