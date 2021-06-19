/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res1 []int
 var res2 []int
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    res1=nil
    res2=nil
    dfs(root1,1)
    dfs(root2,2)
    if len(res1)!=len(res2){
        return false
    }
    for i:=0;i<len(res1);i++{
        if res1[i]!=res2[i]{
            return false
        }
    }
    return true
}

func dfs(root *TreeNode,flag int){
    if root==nil{
        return
    }
    if root.Left==nil&&root.Right==nil{
        //leaf
        if flag==1{
            res1=append(res1,root.Val)
        }else{
            res2=append(res2,root.Val)
        }
    }
    dfs(root.Left,flag)
    dfs(root.Right,flag)
}
