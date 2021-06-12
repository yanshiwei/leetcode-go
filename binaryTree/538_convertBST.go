/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func convertBST(root *TreeNode) *TreeNode {
    /*
    本题中要求我们将每个节点的值修改为原来的节点值加上所有大于它的节点值之和。这样我们只需要反序中序遍历该二叉搜索树，记录过程中的节点值之和，并不断更新当前遍历到的节点的节点值，即可得到题目要求的累加树。
    */
    res=0
    inorderReverse(root)
    return root
}
func inorderReverse(root *TreeNode){
    if root==nil{
        return
    }
    //反序中序遍历,正常是左中右，反序则是右中左
    inorderReverse(root.Right)
    res+=root.Val
    root.Val=res
    inorderReverse(root.Left)
    return
}
