/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var arr []int
    if root==nil {
        return arr
    }
    arr=append(arr,root.Val)
    leftArr:=preorderTraversal(root.Left)
    arr=append(arr,leftArr...)
    rightArr:=preorderTraversal(root.Right)
    arr=append(arr,rightArr...)
    return arr
}
