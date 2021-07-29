/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    if root==nil||p==nil{
        return nil
    }
    //中顺非递归遍历
    var cur=root
    var find=false
    var stack[]*TreeNode
    for len(stack)>0||cur!=nil{
        for cur!=nil{
            stack=append(stack,cur)
            cur=cur.Left
        }
        node:=stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        if find{
            return node
        }
        if node==p{
            find=true
        }
        cur=node.Right
    }
    return nil
}
