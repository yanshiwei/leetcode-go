/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    var length =1
    var cur=head
    for cur!=nil{
        length++
        cur=cur.Next
    }
    cur=head
    var idx=1
    for cur!=nil{
        if idx==length-k{
            return cur.Val
        }
        idx++
        cur=cur.Next
    }
    return 0
}
