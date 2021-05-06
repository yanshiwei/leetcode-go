/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var l int
    var cur=head
    for cur!=nil{
        cur=cur.Next
        l++
    }
    //fmt.Println(l,l/2+1)
    cur=head
    var curLen int
    for cur!=nil{
        curLen++
        if curLen==l/2+1{
            return cur
        }
        cur=cur.Next
    }
    return nil
}
