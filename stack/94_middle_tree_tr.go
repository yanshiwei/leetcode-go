/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var arr []int
    if root==nil {
        return arr
    }
    arrLeft:=inorderTraversal(root.Left)
    arr=append(arr,arrLeft...)
    arr=append(arr,root.Val)//middle
    arrRight:=inorderTraversal(root.Right)
    arr=append(arr,arrRight...)
    return arr
}
