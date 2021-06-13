/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var list []int
func findSecondMinimumValue(root *TreeNode) int {
    list=nil
    dfs(root)
    sort.Ints(list)
    first:=list[0]
    //fmt.Println(list)
    if first==list[len(list)-1]{
        //all the same
        return -1
    }
    for i:=0;i<len(list);i++{
        if first!=list[i]{
            return list[i]
        }
    }
    return -1
}

func dfs(root *TreeNode){
    if root==nil{
        return
    }
    list=append(list,root.Val)
    dfs(root.Left)
    dfs(root.Right)
}
