
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func goodNodes(root *TreeNode) int {
    res=0
    last:=root.Val
    dfs(root,last)
    return res
}
func dfs(root *TreeNode,last int){
    if root==nil{
        return
    }
    if root.Val>=last{
        last=root.Val
        res++
    }
    dfs(root.Left,last)
    dfs(root.Right,last)
}
