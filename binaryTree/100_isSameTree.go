/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return helper(p,q)
}

func helper(p *TreeNode, q *TreeNode)bool{
    if p==nil&&q==nil{
        return true
    }
    if p==nil||q==nil{
        return false
    }
    if p.Val!=q.Val{
        return false
    }
    return helper(p.Left,q.Left)&&helper(p.Right,q.Right)
}
