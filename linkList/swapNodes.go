/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    //在头节点前添加哑节点避免越界且允许了头节点的交换
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    //初始化pre和快慢指针都指向哑节点处
    var pre=dummy
    var slow=dummy
    var fast=dummy
    //pre和快指针走k步，找到正数第k个节点
    for i:=0;i<k;i++{
        pre=pre.Next
        fast=fast.Next
    }
    //pre指向第k个节点
    //快慢指针再一起走。快指针到链表末尾时慢指针指向倒数第k个节点
    for fast!=nil{
        fast=fast.Next
        slow=slow.Next
    }
    //交换两个节点的值
    var tmp=pre.Val
    pre.Val=slow.Val
    slow.Val=tmp
    return head
}
