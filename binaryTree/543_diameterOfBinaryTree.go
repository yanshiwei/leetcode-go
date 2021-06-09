/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int=0
func diameterOfBinaryTree(root *TreeNode) int {
    res=0
    if root==nil{
        return res
    }
    maxPath(root)
    return res
}

func maxPath(root *TreeNode)int{
    if root.Left==nil&&root.Right==nil{
        return 0
    }
    var leftV int
    //判断左子节点是否为空，从而更新左边最长路径
    if root.Left!=nil{
        leftV=maxPath(root.Left)+1
    }else{
        leftV=0
    }
    var RightV int
    if root.Right!=nil{
        RightV=maxPath(root.Right)+1
    }else{
        RightV=0
    }
    res=max(res,leftV+RightV)//全局最长路径
    return max(leftV,RightV)//返回以root点开始能提供的最长路径
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
