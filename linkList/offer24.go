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
    var cur=head
    for cur.Next!=nil{
        cur=cur.Next
    }
    head,_=reverse(head,cur)
    return head
}

func reverse(head,tail *ListNode)(*ListNode,*ListNode){
    var pre=tail.Next
    var cur=head
    for pre!=tail{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    return tail,head
}
