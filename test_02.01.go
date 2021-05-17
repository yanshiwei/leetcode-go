/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    var fre=make(map[int]int)
    var pre=head
    var cur=head
    cur=head
    for cur!=nil{
        fre[cur.Val]++
        next:=cur.Next
        if fre[cur.Val]>1{
            pre.Next=next
        }else{
            //first occurs
            pre=cur
        }
        cur=next
    }
    return head
}
