/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var length int
    var cur=head
    for cur!=nil{
        length++
        cur=cur.Next
    }
    var realN=length-n
    if realN==0{
        return head.Next
    }
    cur=head
    var curLength int
    for cur!=nil{
        curLength++
        if curLength==realN{
            //fmt.Println(curLength,length)
            cur.Next=cur.Next.Next
        }
        cur=cur.Next
    }
    return head
}
