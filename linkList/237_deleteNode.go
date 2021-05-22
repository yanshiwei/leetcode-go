/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    var dummy=&ListNode{
        Val:0,
        Next:node,
    }
    var pre=dummy
    for node!=nil{
        next:=node.Next
        if next==nil{
            pre.Next=nil
            break
        }
        node.Val=next.Val
        pre=node
        node=node.Next
    }
    return
}
