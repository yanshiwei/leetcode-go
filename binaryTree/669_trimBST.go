/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    //当node.val < L，那么修剪后的二叉树出现在节点的右边。否则，我们将会修剪树的两边
    //当node.val > R，那么修剪后的二叉树必定出现在节点的左边。
    if root==nil{
        return nil
    }
    if root.Val<low{
        return trimBST(root.Right,low,high)
    }
    if root.Val>high{
        return trimBST(root.Left,low,high)
    }
    //node.val在[low,high]
    root.Left=trimBST(root.Left,low,high)
    root.Right=trimBST(root.Right,low,high)
    return root
}
