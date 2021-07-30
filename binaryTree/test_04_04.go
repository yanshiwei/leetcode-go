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
    delta:=abs(getHeight(root.Left)-getHeight(root.Right))
    if delta>1{
        return false
    }
    return isBalanced(root.Left)&&isBalanced(root.Right)
}

func getHeight(root *TreeNode)int{
    if root==nil{
        return 0
    }
    return max(getHeight(root.Left),getHeight(root.Right))+1
}

func max(x,y int)int {
    if x<y{
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
