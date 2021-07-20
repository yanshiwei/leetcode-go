/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root==nil{
        return res
    }
    var queue[]*TreeNode
    queue=append(queue,root)
    var height int
    for len(queue)>0{
        curLen:=len(queue)
        var tmp =make([]int,curLen)
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            if height%2==0{
                tmp[i]=cur.Val
            }else{
                tmp[curLen-1-i]=cur.Val
            }
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        height++
        res=append(res,tmp)
    }
    return res
}
