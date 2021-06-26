/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func sumRootToLeaf(root *TreeNode) int {
    res=0
    dfs(root,0)
    return res
}

func dfs(root *TreeNode,sum int){
    if root==nil{
        return
    }
    if root.Left==nil&&root.Right==nil{
        sum=2*sum+root.Val
        //fmt.Println(root.Val,sum)
        res+=sum
    }else{
        sum=2*sum+root.Val
    }
    dfs(root.Left,sum)
    dfs(root.Right,sum)
}
