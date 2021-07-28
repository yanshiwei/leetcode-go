/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root==nil{
        return root
    }
    if root==p || root==q {
        return root
    }
    left:=lowestCommonAncestor(root.Left,p,q)
    right:=lowestCommonAncestor(root.Right,p,q)
    // p or q in right sub tree
    if left==nil{
        return right
    }
    // p or q in left
    if right==nil{
        return left
    }
    //p q in left and right
    return root
}
