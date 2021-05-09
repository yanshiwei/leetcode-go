/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    //创建哑结点 dummyHead，令 dummyHead.next = head
    var dummyHead=&ListNode{
        Val:0,
        Next:head,
    }
    //令 cur 表示当前到达的节点，初始时 cur = dummyHead。每次需要交换 cur 后面的两个节点。
    var cur=dummyHead
    //如果 cur 的后面没有节点或者只有一个节点，则没有更多的节点需要交换，因此结束交换
    for cur.Next!=nil&&cur.Next.Next!=nil{
        node1:=cur.Next
        node2:=cur.Next.Next
        cur.Next=node2
        node1.Next=node2.Next
        node2.Next=node1
        cur=node1
    }
    return dummyHead.Next
}
