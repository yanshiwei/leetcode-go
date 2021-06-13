/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var fre map[int]int
func findTarget(root *TreeNode, k int) bool {
    fre=make(map[int]int)
    return dfs(root,k) 
}

func dfs(root *TreeNode, k int)bool{
    if root==nil{
        return false
    }
    if _,ok:=fre[k-root.Val];ok{
        return true
    }
    fre[root.Val]=1
    return dfs(root.Left,k)||dfs(root.Right,k)
}
