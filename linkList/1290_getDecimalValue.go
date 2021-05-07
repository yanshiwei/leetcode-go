/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    var res int
    var cur=head
    for cur!=nil{
        res=res*2+cur.Val
        cur=cur.Next
    }
    return res
}
