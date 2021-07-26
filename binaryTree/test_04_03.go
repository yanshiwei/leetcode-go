/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    if tree==nil{
        return nil
    }
    var res []*ListNode
    var queue[]*TreeNode
    queue=append(queue,tree)
    for len(queue)>0{
        //each level
        curLen:=len(queue)
        dummy:=&ListNode{}
        cur:=dummy
        for i:=0;i<curLen;i++{
            node:=queue[0]
            queue=queue[1:]
            cur.Next=&ListNode{Val:node.Val}
            cur=cur.Next
            if node.Left!=nil{
                queue=append(queue,node.Left)
            }
            if node.Right!=nil{
                queue=append(queue,node.Right)
            }
        }
        res=append(res,dummy.Next)
    }
    return res
}
