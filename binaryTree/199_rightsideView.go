/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var res []int
    if root==nil{
        return res
    }
    var queue[]*TreeNode
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        res=append(res,queue[curLen-1].Val)
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
    }
    return res
}
