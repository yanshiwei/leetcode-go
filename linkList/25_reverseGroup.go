/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    //参考：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/#comment
    //头节点 head 前面是没有节点 pre，构造一个哑节点
    var dummyNode=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummyNode
    var cur=head
    for cur!=nil{
        var tail =pre
        //查看剩余部分长度是否大于等于 k
        for i:=0;i<k;i++{
            tail=tail.Next
            if tail==nil{
                return dummyNode.Next
            }
        }
        //获取长k的子链表from cur to tail
        next:=tail.Next
        cur,tail=reverse(cur,tail)
        //把子链表重新接回原链表
        pre.Next=cur
        tail.Next=next
        //从下一个k段开始
        pre=tail
        cur=tail.Next
    }
    return dummyNode.Next
}

func reverse(head,tail *ListNode)(*ListNode,*ListNode){
    //翻转一个子链表，并且返回新的头与尾
    var pre=tail.Next
    var cur=head
    for pre!=tail{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    //fmt.Println(tail.Val,head.Val)
    return tail,head
}
