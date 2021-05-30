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
    //遍历左 root 右 所有情况
    return delta<=1&&isBalanced(root.Left)&&isBalanced(root.Right)
}
func max(x,y int)int{
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
//判断一个根节点上root的子树的高度
func getHeight(root *TreeNode)int{
    if root==nil{
        return 0
    }
    return max(getHeight(root.Left),getHeight(root.Right))+1
}
