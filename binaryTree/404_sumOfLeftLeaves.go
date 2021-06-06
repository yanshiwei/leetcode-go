/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    //左叶子+非右叶子
    if root==nil{
        return 0
    }
    res:=helper(root)
    return res
}

func helper(root *TreeNode)int{
    var res int
    if root.Left!=nil{
        if root.Left.Left==nil&&root.Left.Right==nil{
            //左叶子
            res+=root.Left.Val
        }else{
            res+=helper(root.Left)
        }
    }
    if root.Right!=nil{
        if root.Right.Left!=nil||root.Right.Right!=nil{
            //非右叶子
            res+=helper(root.Right)
        }
    }
    return res
}
