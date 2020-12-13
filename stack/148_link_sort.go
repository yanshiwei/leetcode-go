/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
 分而治之，先将链表分成不能再分为止（保证最简单的有序），每次分法采用对半，通过快慢指针实现
 合并两个的有序链表
 */
 func merge(l1 *ListNode, l2 *ListNode) *ListNode {
     var dummy=new(ListNode)
     var cur=dummy
     for l1!=nil && l2!=nil {
         if l1.Val<l2.Val {
             cur.Next=l1
             l1=l1.Next
         }else {
             cur.Next=l2
             l2=l2.Next
         }
         cur=cur.Next
     }
     if l1!=nil {cur.Next=l1}
     if l2!=nil {cur.Next=l2}
     return dummy.Next
 }

func halfList(head *ListNode)(*ListNode, *ListNode){
	if head == nil{
		return nil, nil
	}
	fast := head.Next
	slow := head
	for fast!= nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	slowNext := slow.Next
	slow.Next = nil
	return head, slowNext
}

func sortList(head *ListNode) *ListNode {
    if head==nil || head.Next==nil {
        return head
    }
    left,right:=halfList(head)
	return merge(sortList(left), sortList(right))
}
