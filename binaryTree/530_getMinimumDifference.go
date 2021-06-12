/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int
var base int
var delta int
func getMinimumDifference(root *TreeNode) int {
    base=-1
    delta=INTMAX
    inorder(root)
    return delta
}

func inorder(root *TreeNode){
    if root==nil{
        return
    }
    inorder(root.Left)
    updateRes(root.Val)
    inorder(root.Right)
}

func updateRes(x int){
    if base==-1{
        //init
        base=x
    }else{
        //fmt.Println(x,base)
        curDelta:=abs(x-base)
        delta=min(delta,curDelta)
        base=x
    }
}

func abs(x int)int {
    if x<0{
        return -1*x
    }
    return x
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
const INTMAX=int(^uint(0)>>1)
