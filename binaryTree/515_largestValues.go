/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    var res []int
    if root==nil{
        return res
    }
    var queue=make([]*TreeNode,0)
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        var curMax=INTMIN
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            if curMax<cur.Val{
                curMax=cur.Val
            }
            queue=queue[1:]
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        res=append(res,curMax)
    }
    return res
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
