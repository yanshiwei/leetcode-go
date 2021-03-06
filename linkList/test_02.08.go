/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    /*
   有环单链表的环连接点位置
   head->joint:a
   joint->pos:b
   pos->joint:c
   设链表中环外部分的长度为 a。slow 指针进入环后，又走了 b的距离在pos处与 fast 相遇。此时，fast 指针已经走完了环的 n 圈，因此它走过的总距离为 a+n(b+c)+b=a+(n+1)b+nc而slow走的距离是a+b
   又因为fast 指针走过的距离都为slow 指针的 2 倍
   a+(n+1)b+nc=2*（a+b)=>a=c+(n-1)(b+c)
   即从相遇点pos到入环点joint的距离c加上 n-1圈的环长，恰好等于从链表头部到入环点的距离a
   则从相遇点pos开始，链表头部开始走和从pos开始走，每次向后移动一个位置。最终，它们会在入环点相遇。
    */
    if head==nil||head.Next==nil{
        return nil
    }
    var slow=head
    var fast=head
    for fast!=nil&&fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
        if slow==fast{
            //第一次相遇
            var idx=head
            for idx!=slow{
                idx=idx.Next
                slow=slow.Next
            }
            return idx
        }
    }
    return nil
}
