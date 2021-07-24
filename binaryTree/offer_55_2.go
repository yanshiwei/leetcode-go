/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root==nil{
        return true
    }
    delat:=abs(height(root.Left)-height(root.Right))
    if delat>1{
        return false
    }
    return isBalanced(root.Left)&&isBalanced(root.Right)
}

func height(root *TreeNode)int {
    if root==nil{
        return 0
    }
    return max(height(root.Left),height(root.Right))+1
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

func abs(x int)int {
    if x<0{
        return -1*x
    }
    return x
}
