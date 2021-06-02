/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return helper(root,0)
}

func helper(root *TreeNode,num int)int {
    if root==nil{
        return 0
    }
    tmp:=num*10+root.Val
    if root.Left==nil&&root.Right==nil{
        //leave
        return tmp
    }
    return helper(root.Left,tmp)+helper(root.Right,tmp)
}
