/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var base int
func isUnivalTree(root *TreeNode) bool {
    base=0
    return dfs(root)
}

func dfs(root *TreeNode)bool{
    if root==nil{
        return true
    }
    //fmt.Println(root.Val)
    if base==0{
        base=root.Val
    }else{
        if base!=root.Val{
            return false
        }
    }
    return dfs(root.Left)&&dfs(root.Right)
}
