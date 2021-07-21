/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
var idx int
func kthLargest(root *TreeNode, k int) int {
    res=0
    idx=k
    dfs(root)
    return res
}

func dfs(root *TreeNode){
    if root==nil{
        return
    }
    dfs(root.Right)
    idx--
    if idx==0 {
        res=root.Val
        return
    }
    dfs(root.Left)
}
