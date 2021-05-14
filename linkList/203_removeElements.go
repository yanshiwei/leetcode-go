/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head==nil{
        return nil
    }
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummy
    var cur=head
    for cur!=nil{
        next:=cur.Next
        if cur.Val==val{
            pre.Next=next
        }else{
            pre=cur
        }
        cur=next
    }
    return dummy.Next
}
