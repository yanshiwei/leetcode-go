/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root==nil{
        return true
    }
    return dfs(root,INTMIN,INTMAX)
}

func dfs(root *TreeNode,low,high int)bool{
    if root== nil{
        return true
    }
    if root.Val<=low||root.Val>=high{
        return false
    }
    return dfs(root.Left,low,root.Val)&&dfs(root.Right,root.Val,high)
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
