/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var fre=make(map[int]int)
    var cur =head
    for cur!=nil{
        fre[cur.Val]++
        cur=cur.Next
    }
    cur =head
    //头节点没有pre,构造一个
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummy
    for cur!=nil{
        next:=cur.Next
        if fre[cur.Val]>1{
            pre.Next=next
            //要删除cur，则pre仍然不变
        }else{
            pre=cur
        }
        cur=next
    }
    return dummy.Next
}
