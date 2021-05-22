/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    var length int
    var cur=head
    for cur!=nil{
        length++
        cur=cur.Next
    }
    var res=make([]int,length)
    cur=head
    var idx int
    for cur!=nil{
        res[length-1-idx]=cur.Val
        idx++
        cur=cur.Next
    }
    return res
}
