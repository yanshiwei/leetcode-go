/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverse(a[]int)[]int {
    var arr []int
    for i:=len(a)-1;i>=0;i-- {
        arr=append(arr,a[i])
    }
    return arr
}
func zigzagLevelOrder(root *TreeNode) [][]int {
    var arr [][]int
    if root==nil {
        return arr
    }
    //level traval by queue
    var queue []*TreeNode
    var level uint32
    queue=append(queue,root)
    for len(queue)>0 {
        var cnt=len(queue)
        var one []int
        for i:=0;i<cnt;i++ {
            var tmp=queue[0]
            queue=queue[1:]//pop
            if tmp!=nil {
                one=append(one,tmp.Val)
                if tmp.Left!=nil {
                    queue=append(queue,tmp.Left)//push
                }
                if tmp.Right!=nil {
                    queue=append(queue,tmp.Right)//push
                }
            }
        }
        if level%2==0 {
            arr=append(arr,one)
        }else {
            arr=append(arr,reverse(one))
        }
        
        level+=1
    }
    return arr
}
