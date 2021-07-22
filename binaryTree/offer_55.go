/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func maxDepth(root *TreeNode) int {
    res=0
    if root==nil{
        return res
    }
    var queue[]*TreeNode
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        res++
    }
    return res
}

