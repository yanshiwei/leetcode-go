/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    //最大二叉树插入一个数后仍然是最大二叉树
    if val>root.Val{
        //新的最大值
        return &TreeNode{Val:val,Left:root}
    }
    //新插入的肯定在右，之前插入的在左
    var parent=root
    var cur=root.Right
    for cur!=nil&&val<cur.Val{
        parent=cur
        cur=cur.Right
    }
    parent.Right=&TreeNode{Val:val,Left:cur}
    return root
}
