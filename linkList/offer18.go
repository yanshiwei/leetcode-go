/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    var cur=head
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummy
    for cur!=nil{
        if cur.Val==val{
            pre.Next=cur.Next
        }
        pre=cur
        cur=cur.Next
    }
    return dummy.Next
}
