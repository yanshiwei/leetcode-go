/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    if root==nil{
        return 0
    }
    var queue []*TreeNode
    var level int
    var resV=INTMIN
    var resI int
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        level++
        var tmp int
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            tmp+=cur.Val
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        if resV<tmp{
            resV=tmp
            resI=level
        }
    }
    return resI
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
