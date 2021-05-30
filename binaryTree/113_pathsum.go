/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res [][]int
var path[]int
func pathSum(root *TreeNode, targetSum int) [][]int {
    res=nil
    path=nil
    dfs(root,targetSum)
    return res
}

func dfs(root *TreeNode, targetSum int){
    if root==nil{
        return 
    }
    path=append(path,root.Val)
    targetSum-=root.Val
    if root.Left==nil&&root.Right==nil{
        //fmt.Println(root.Val,path,targetSum)
        if targetSum==0{
            var finalPath []int
            finalPath=append([]int(nil),path...)
            res=append(res,finalPath)
        }
    }
    dfs(root.Left,targetSum)
    dfs(root.Right,targetSum)
    path=path[0:len(path)-1]//pop back
    return
}
