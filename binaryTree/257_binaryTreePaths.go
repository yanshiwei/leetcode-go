/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []string
var path []int
func binaryTreePaths(root *TreeNode) []string {
    //自顶向下，DFS前序遍历
    res=nil
    path=nil
    helper(root)
    return res
}
func helper(root *TreeNode){
    if root==nil{
        return 
    }
    path=append(path,root.Val)
    if root.Left==nil&&root.Right==nil{
        var finalRes string
        for i:=0;i<len(path)-1;i++{
            finalRes+=strconv.Itoa(path[i])+"->"
        }
        finalRes+=strconv.Itoa(path[len(path)-1])
        res=append(res,finalRes)
    }
    helper(root.Left)
    helper(root.Right)
    path=path[0:len(path)-1]//pop back
    return
}
