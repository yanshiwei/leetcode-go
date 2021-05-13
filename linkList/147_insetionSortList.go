/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    var dummy=&ListNode{
        Val:0,
        Next:nil,
    }
    var pre=dummy
    var cur=head
    for cur!=nil{
        pre=dummy
        next:=cur.Next
        //寻找插入位置，pre之后插入
        for pre.Next!=nil&&pre.Next.Val<cur.Val{
            pre=pre.Next
        }
        cur.Next=pre.Next
        pre.Next=cur
        cur=next
    }
    return dummy.Next
}
