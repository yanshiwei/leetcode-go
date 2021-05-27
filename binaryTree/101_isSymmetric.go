/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    /*
    如果同时满足下面的条件，两个树互为镜像：
    1.它们的两个根结点具有相同的值
    2.每个树的右子树都与另一个树的左子树镜像对称
    类似两个树是否相等
    */
    return helper(root,root)
}

func helper(p,q *TreeNode)bool{
    if p==nil&&q==nil{
        return true
    }
    if p==nil||q==nil{
        return false
    }
    if p.Val!=q.Val{
        return false
    }
    return helper(p.Left,q.Right)&&helper(p.Right,q.Left)
}
