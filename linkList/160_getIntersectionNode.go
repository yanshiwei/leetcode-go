/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var fre=make(map[*ListNode]struct{})
    var cur=headA
    for cur!=nil{
        fre[cur]=struct{}{}
        cur=cur.Next
    }
    cur=headB
    for cur!=nil{
        if _,ok:=fre[cur];ok{
            return cur
        }
        cur=cur.Next
    }
    return nil
}
