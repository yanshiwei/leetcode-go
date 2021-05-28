/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    var res [][]int
    if root==nil{
        return res
    }
    var queue []*TreeNode
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        var tmp[]int
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            tmp=append(tmp,cur.Val)
            if cur.Left!=nil{
            queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        res=append(res,tmp)
        if len(res)>1{
            copy(res[1:],res[0:len(res)-1])
            res[0]=tmp
        }
    }
    return res
}
