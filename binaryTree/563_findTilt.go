/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func findTilt(root *TreeNode) int {
    res=0
    postorder(root)
    return res
}
//root节点可以提供的子树节点和
func postorder(root *TreeNode)int{
    if root==nil{
        return 0
    }
    left:=postorder(root.Left)
    right:=postorder(root.Right)
    cur:=abs(left-right)//root节点的左右子树的坡度，更新全局变量
    res+=cur
    return root.Val+left+right//root节点可以提供的子树节点和
}

func abs(x int)int {
    if x<0{
        return -1*x
    }
    return x
}
