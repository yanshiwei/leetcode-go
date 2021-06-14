/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func bstToGst(root *TreeNode) *TreeNode {
    res=0
    inorderReverse(root)
    return root
}

func inorderReverse(root*TreeNode) *TreeNode{
    //反序中序遍历
    if root==nil{
        return nil
    }
    inorderReverse(root.Right)
    res+=root.Val
    root.Val=res
    inorderReverse(root.Left)
    return root
}
