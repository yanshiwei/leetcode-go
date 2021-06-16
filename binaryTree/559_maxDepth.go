/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func maxDepth(root *Node) int {
    if root==nil{
        return 0
    }else if len(root.Children)==0{
        return 1
    }else{
        var queue []*Node
        var res int
        queue=append(queue,root)
        for len(queue)>0{
            curLen:=len(queue)
            res++
            for i:=0;i<curLen;i++{
                cur:=queue[0]
                queue=queue[1:]
                for j:=0;j<len(cur.Children);j++{
                    if cur.Children[j]!=nil{
                        queue=append(queue,cur.Children[j])
                    }
                }
            }
        }
        return res
    }
}

