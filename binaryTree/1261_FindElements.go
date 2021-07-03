/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var Fre map[int]bool
type FindElements struct {
}


func Constructor(root *TreeNode) FindElements {
    var m =FindElements{}
    if root==nil{
        return m
    }
    Fre=make(map[int]bool)
    root.Val=0
    Fre[root.Val]=true
    dfs(root)
    return m
}

func dfs(root *TreeNode){
    if root.Left!=nil{
        root.Left.Val=2*root.Val+1
        Fre[root.Left.Val]=true
        dfs(root.Left)
    }
    if root.Right!=nil{
        root.Right.Val=2*root.Val+2
        Fre[root.Right.Val]=true
        dfs(root.Right)
    }
}

func (this *FindElements) Find(target int) bool {
    _,ok:=Fre[target]
    if ok{
        return true
    }
    return false
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
