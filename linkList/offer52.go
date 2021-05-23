/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var curA=headA
    var curB=headB
    for curA!=curB{
        if curA!=nil{
            curA=curA.Next
        }else{
            curA=headB
        }
        if curB!=nil{
            curB=curB.Next
        }else{
            curB=headA
        }
    }
    return curB
}
