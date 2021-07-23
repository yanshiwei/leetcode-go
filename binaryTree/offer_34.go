/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res [][]int
 var path []int
func pathSum(root *TreeNode, target int) [][]int {
    res=nil
    path=nil
    dfs(root,target)
    return res
}
func dfs(root *TreeNode, target int){
    if root==nil{
        return
    }
    //pre
    path=append(path,root.Val)//push
    target-=root.Val
    if root.Left==nil&&root.Right==nil{
        //leaf
        if target==0{
            var tmp []int
            tmp=append([]int(nil),path...)
            res=append(res,tmp)
        }
    }
    dfs(root.Left,target)
    dfs(root.Right,target)
    //recover ctx
    path=path[0:len(path)-1]//pop
    return
}

