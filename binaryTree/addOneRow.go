/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth==1{
        //根
        tmp:=&TreeNode{Val:val}
        tmp.Left=root
        return tmp
    }else{
        //其它层
        var queue []*TreeNode
        var curHeight int
        queue=append(queue,root)
        for len(queue)>0{
            curLen:=len(queue)
            curHeight++
            for i:=0;i<curLen;i++{
                cur:=queue[0]
                queue=queue[1:]
                if curHeight==depth-1{
                    //cur的新的左节点
                    newL:=&TreeNode{Val:val}
                    oldL:=cur.Left
                    cur.Left=newL
                    cur.Left.Left=oldL
                    //cur的新的右节点
                    newR:=&TreeNode{Val:val}
                    oldR:=cur.Right
                    cur.Right=newR
                    cur.Right.Right=oldR
                }
                if cur.Left!=nil{
                    queue=append(queue,cur.Left)
                }
                if cur.Right!=nil{
                    queue=append(queue,cur.Right)
                }
            }
        }
        return root
    }
}
