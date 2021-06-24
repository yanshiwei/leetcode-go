/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var queue []Info
    queue=append(queue,Info{CurNode:root,PaNode:nil})
    for len(queue)>0{
        curLen:=len(queue)
        var tmpPa []*TreeNode
        for i:=0;i<curLen;i++{
            info:=queue[0]
            queue=queue[1:]
            node:=info.CurNode
            parent:=info.PaNode
            if node.Val==x||node.Val==y{
                tmpPa=append(tmpPa,parent)
            }
            if node.Left!=nil{
                queue=append(queue,Info{CurNode:node.Left,PaNode:node})
            }
            if node.Right!=nil{
                queue=append(queue,Info{CurNode:node.Right,PaNode:node})
            }
        }
        // `x` 和 `y` 都没出现
        if len(tmpPa)==0{
            continue
        }else if len(tmpPa)==1{
            // `x` 和 `y` 只出现一个
            return false
        }else if len(tmpPa)==2{
             // `x` 和 `y` 都出现了
             // `x` 和 `y` 父节点 相同/不相同 ？
             if tmpPa[0]==tmpPa[1]{
                 return false
             }
             return true
        }
    }
    return false
}

type Info struct{
    CurNode *TreeNode
    PaNode *TreeNode
}
