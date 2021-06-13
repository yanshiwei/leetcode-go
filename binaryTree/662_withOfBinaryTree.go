/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func widthOfBinaryTree(root *TreeNode) int {
    //满二叉树的第i层至多有2^(i-1)个结点（i>=1）
    //左子树是父节点的index * 2,右子树是 index * 2 + 1,减去start * 2 便是在该层的位置，用于防止索引太大溢出
    res=INTMIN
    var queue []Info
    queue=append(queue,Info{Node:root,Idx:1})//1开始
    var curLevel int
    for len(queue)>0{
        curLen:=len(queue)
        curLevel++
        start:=queue[0].Idx//start是本层起点
        var idx int
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            node:=cur.Node
            idx=cur.Idx
            if node.Left!=nil{
                queue=append(queue,Info{Node:node.Left,Idx:idx*2-start*2})
            }
            if node.Right!=nil{
                queue=append(queue,Info{Node:node.Right,Idx:idx*2+1-start*2})
            }
        }
        res=max(res,idx-start+1)
    }
    return res
}
type Info struct{
    Node *TreeNode
    Idx int
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
