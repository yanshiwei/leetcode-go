/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    var cur=head
    var length =1
    for cur!=nil{
        length++
        cur=cur.Next
    }
    //倒数k就是正数length-k
    cur=head
    var idx=1
    for cur!=nil{
        if idx==length-k{
            return cur
        }
        idx++
        cur=cur.Next
    }
    return nil
}
