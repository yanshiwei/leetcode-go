/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1==nil&&root2==nil{
        return nil
    }
    if root1!=nil&&root2==nil{
        return root1
    }
    if root1==nil&&root2!=nil{
        return root2
    }
    tmp:=&TreeNode{Val:root1.Val+root2.Val}
    tmp.Left=mergeTrees(root1.Left,root2.Left)
    tmp.Right=mergeTrees(root1.Right,root2.Right)
    return tmp
}
