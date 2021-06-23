/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    var queue[]Info
    queue=append(queue,Info{Node:root,Idx:1})
    var i int//depth 
    for i<len(queue){
        cur:=queue[i]
        node:=cur.Node
        if node!=nil{
            queue=append(queue,Info{Node:node.Left,Idx:cur.Idx*2})
            queue=append(queue,Info{Node:node.Right,Idx:cur.Idx*2+1})
        }
        i++
    }
    //fmt.Println(queue,i)
    //检测编号序列是否为无间隔.树中所有节点的编号按照广度优先搜索顺序正好是升序,只需要检查最后一个编号是否正确，因为最后一个编号的值最大. i=1,2,3,4,5,6,7,8,...2^(h)-1
    if queue[i-1].Idx==len(queue){
        return true
    }
    return false
}

type Info struct{
    Node *TreeNode
    Idx int
}
