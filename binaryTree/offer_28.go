/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root==nil{
        return true
    }
    return dfs(root.Left,root.Right)
}

func dfs(left,right *TreeNode)bool {
    if left==nil&&right==nil{
        return true
    }
    if left==nil||right==nil||left.Val!=right.Val{
        return false
    }
    //节点 L.left 和 R.right 是否对称；L.right 和R.left 是否对称
    return dfs(left.Left,right.Right)&&dfs(left.Right,right.Left)
}
