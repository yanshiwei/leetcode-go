/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func sumEvenGrandparent(root *TreeNode) int {
    res=0
    dfs(root,nil,nil)
    return res
}

func dfs(root, parent,grandParent *TreeNode){
    if root==nil{
        return
    }
    if grandParent!=nil&&grandParent.Val%2==0{
        res+=root.Val
    }
    dfs(root.Left,root,parent)
    dfs(root.Right,root,parent)
}
