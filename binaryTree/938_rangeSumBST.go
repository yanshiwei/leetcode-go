/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func rangeSumBST(root *TreeNode, low int, high int) int {
    res=0
    dfs(root,low,high)
    return res
}

func dfs(root *TreeNode, low int, high int){
    if root==nil{
        return
    }
    dfs(root.Left,low,high)
    if root.Val>=low&&root.Val<=high{
        res+=root.Val
    }
    dfs(root.Right,low,high)
}
