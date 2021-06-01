/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }
    var queue []*TreeNode
    var minDepth int=1// 1 is root
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        //遍历该层
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            if cur.Left==nil&&cur.Right==nil{
                //leaf
                return minDepth
            }
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        minDepth++
    }
    return minDepth
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
