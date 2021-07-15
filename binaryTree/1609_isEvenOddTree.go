/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    if root==nil{
        return true
    }
    var queue []*TreeNode
    var idx int
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        var last int
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            if idx%2==0{
                //偶数层
                if cur.Val%2==0{
                    return false
                }
                if last==0{
                    last=cur.Val
                }else{
                    if cur.Val<=last{
                        return false
                    }
                }
            }else{
                //奇数层
                if cur.Val%2==1{
                    return false
                }
                if last==0{
                    last=cur.Val
                }else{
                    if cur.Val>=last{
                        return false
                    }
                }               
            }
            last=cur.Val
            if cur.Left!=nil{
                queue=append(queue,cur.Left)
            }
            if cur.Right!=nil{
                queue=append(queue,cur.Right)
            }
        }
        idx++
    }
    return true
}
