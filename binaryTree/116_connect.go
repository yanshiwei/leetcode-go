/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root==nil{
        return root
    }
    var queue []*Node
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        // 遍历这一层的所有节点
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            if i<curLen-1{
                //该层的非最后一个节点
                cur.Next=queue[0]
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
