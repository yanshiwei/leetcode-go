/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    /*
    只需维护两个链表small 和 large 即可，small 链表按顺序存储所有小于 x 的节点，large 链表按顺序存储所有大于等于 x 的节点。遍历完原链表后，我们只要将 small 链表尾节点指向 large 链表的头节点即能完成对链表的分隔
    我们设smallDummy 和largeDummy分别为两个链表的哑节点，即它们的 next 指针指向链表的头节点，这样做的目的是为了更方便地处理头节点为空的边界条件。同时设 smallCur和 largeCur 节点指向当前链表的末尾节点
    */
    var smallCur=&ListNode{
        Val:0,
        Next:nil,
    }
    var smallDummy=smallCur
    var largeCur=&ListNode{
        Val:0,
        Next:nil,
    }
    var largeDummy=largeCur
    for head!=nil{
        if head.Val<x{
            smallCur.Next=head
            smallCur=smallCur.Next
        }else{
            largeCur.Next=head
            largeCur=largeCur.Next
        }
        head=head.Next
    }
    //连接两个链表
    largeCur.Next=nil
    smallCur.Next=largeDummy.Next
    return smallDummy.Next
}

