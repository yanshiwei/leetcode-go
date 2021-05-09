/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var fre=make([]int,201)
    var cur=head
    for cur!=nil{
        fre[cur.Val+100]++
        cur=cur.Next
    }
    //头节点没有pre，构造一个
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummy
    cur=head
    for cur!=nil{
        next:=cur.Next
        if fre[cur.Val+100]>1{
            pre.Next=next
            fre[cur.Val+100]--
        }else{
            pre=cur
        }
        cur=next
    }
    return dummy.Next
}
