/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
	return 0
   }

   left := maxheight(root.Left)
   right := maxheight(root.Right)

     diameter := left + right

    sub := max(diameterOfBinaryTree(root.Left),
              diameterOfBinaryTree(root.Right))

    return max(diameter, sub)
}



func maxheight(root *TreeNode) int {
   if root == nil {
	return 0
   }

   return 1 + max(maxheight(root.Left),maxheight(root.Right))

}