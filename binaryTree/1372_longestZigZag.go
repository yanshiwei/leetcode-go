/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func longestZigZag(root *TreeNode) int {
    /*
    引入辅助函数 helper，该函数保存当前长度 res 和 是否为左节点的 isLeft，逻辑为：
1.如果是左节点，在向下递归时，右节点 length+1，isLeft 为 false；左节点长度初始化为1
2.如果是右节点，在向下递归时，左节点 length+1，isLeft 为 false；右节点长度初始化为1
3.遇到空节点时递归终止
    递归时遇到更长的长度，更新 max
    */
    res=0
    if root==nil{
        return 0
    }
    helper(root,0,true)
    helper(root,0,false)
    return res
}

func helper(root*TreeNode,length int,isLeft bool){
    if root==nil{
        return

    }
    if isLeft{
        //left node
        helper(root.Right,length+1,false)
        helper(root.Left,1,true)//重新开始
    }else{
        //right node
        helper(root.Left,length+1,true)
        helper(root.Right,1,false)//重新开始
    }
    if length>res{
        res=length
    }
}
