/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res []*TreeNode
func increasingBST(root *TreeNode) *TreeNode {
    res=nil
    inorder(root)
    dummy:=&TreeNode{
        Val:-1,
    }
    curNode:=dummy
    for i:=0;i<len(res);i++{
        res[i].Left=nil
        curNode.Right=res[i]
        curNode=curNode.Right
    }
    return dummy.Right
}
func inorder(root *TreeNode){
    if root==nil{
        return 
    }
    inorder(root.Left)
    res=append(res,root)
    inorder(root.Right)
}
