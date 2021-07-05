/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res []int
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    res=nil
    dfs(root1)
    //fmt.Println(res)
    dfs(root2)
    //fmt.Println(res)
    sort.Ints(res)
    return res
}

func dfs(root *TreeNode){
    if root==nil{
        return
    }
    res=append(res,root.Val)
    dfs(root.Left)
    dfs(root.Right)
}
