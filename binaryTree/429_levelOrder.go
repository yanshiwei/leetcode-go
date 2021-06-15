/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    var res [][]int
    if root==nil{
        return res
    }
    var queue []*Node
    queue=append(queue,root)
    for len(queue)>0{
        curLen:=len(queue)
        var curList []int
        for i:=0;i<curLen;i++{
            cur:=queue[0]
            queue=queue[1:]
            curList=append(curList,cur.Val)
            for j:=0;j<len(cur.Children);j++{
                if cur.Children[j]!=nil{
                    queue=append(queue,cur.Children[j])
                }
            }
        }
        res=append(res,curList)
    }
    return res
}
