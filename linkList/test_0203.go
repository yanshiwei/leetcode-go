/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    if node==nil{
        return
    }
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
        pre=node
        node.Val=next.Val
        node=next
    }
    return
}
