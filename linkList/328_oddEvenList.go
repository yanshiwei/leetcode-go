/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head==nil||head.Next==nil||head.Next.Next==nil{
        return head
    }
    //可以将奇数节点和偶数节点分离成奇数链表和偶数链表，然后将偶数链表连接在奇数链表之后，合并后的链表即为结果链表
    var odd=head
    var evenHead=head.Next
    var even=evenHead
    for even!=nil&&even.Next!=nil{
        //奇数的下一个是偶数的下一个
        odd.Next=even.Next
        odd=odd.Next
        //偶数的下一个数奇数的下一个
        even.Next=odd.Next
        even=even.Next
    }
    //连接奇数和偶数链表，偶数在后面
    odd.Next=evenHead
    return head
}
