/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    var res int
    if root==nil{
        return res
    }
    var queue=make([]*TreeNode,0)
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            if i==0{
                res=cur.Val
            }
            queue=queue[1:]
            if cur.Left!=nil{
            queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
    }
    return res
}
