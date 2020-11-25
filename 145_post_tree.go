/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var arr []int
    if root==nil {
        return arr
    }
    leftArr:=postorderTraversal(root.Left)
    arr=append(arr,leftArr...)
    rightArr:=postorderTraversal(root.Right)
    arr=append(arr,rightArr...)
    arr=append(arr,root.Val)
    return arr
}
