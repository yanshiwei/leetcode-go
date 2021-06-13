/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    var queue []*TreeNode
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        var curSum float64
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            curSum+=float64(cur.Val)
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        res=append(res,curSum/float64(curLen))
    }
    return res
}
