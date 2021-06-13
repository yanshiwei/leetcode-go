/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res [][]string
func printTree(root *TreeNode) [][]string {
    height:=getHeight(root)
    //数组m=height,n最大为2^height-1
    res=make([][]string,height)
    for i:=range res{
        res[i]=make([]string,(1<<height)-1)
    }
    dfs(root,0,0,len(res[0]))
    //format res
    return res
}

func getHeight(root *TreeNode)int{
    if root==nil{
        return 0
    }
    return 1+max(getHeight(root.Left),getHeight(root.Right))
}
func max(x, y int)int {
    if x<y{
        return y
    }
    return x
}
//dfs preorder
func dfs(root *TreeNode,curLevel,l,r int){
    if root==nil{
        return
    }
    res[curLevel][l+(r-l)/2]=strconv.Itoa(root.Val)
    dfs(root.Left,curLevel+1,l,l+(r-l)/2)
    dfs(root.Right,curLevel+1,l+(r-l)/2+1,r)
    return
}
