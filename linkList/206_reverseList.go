/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    var pre *ListNode
    var cur=head
    for cur.Next!=nil{
        cur=cur.Next
    }
    var tail=cur
    cur=head
    pre=tail.Next
    for tail!=pre{
        next:=cur.Next
        cur.Next=pre
        //next
        pre=cur
        cur=next
    }
    return tail
}
