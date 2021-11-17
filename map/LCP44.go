/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var fre map[int]struct{}
func numColor(root *TreeNode) int {
    if root==nil{
        return 0
    }
    fre=make(map[int]struct{})
    dfs(root)
    return len(fre)
}
func dfs(root *TreeNode){
    if root==nil{
        return
    }
    fre[root.Val]=struct{}{}
    dfs(root.Left)
    dfs(root.Right)
}
