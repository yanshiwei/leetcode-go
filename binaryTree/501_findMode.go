/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []int
var base,count,maxCnt int
func findMode(root *TreeNode) []int {
    res=nil
    base=0
    count=0
    maxCnt=0
    inOrder(root)
    return res
}

func inOrder(root *TreeNode){
    if root==nil{
        return
    }
    inOrder(root.Left)
    updateRes(root.Val)
    inOrder(root.Right)
}

func updateRes(x int){
    if x==base{
        //有连续
        count++
    }else{
        //第一次出现
        base=x
        count=1
    }
    if count==maxCnt{
        //说明当前的这个数字（base）出现的次数等于当前众数出现的次数
        res=append(res,base)
    }
    if count>maxCnt{
        //那么说明当前的这个数字（base）出现的次数大于当前众数出现的次数
        maxCnt=count
        res=[]int{base}
    }
}
