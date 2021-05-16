/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    var firstPre,secondCur *ListNode
    var cur=list1
    var dummy=&ListNode{
        Val:0,
        Next:list1,
    }
    var pre=dummy
    var idx int
    for cur!=nil{
        if idx==a{
            firstPre=pre
        }
        if idx==b{
            secondCur=cur
        }
        idx++
        pre=cur
        cur=cur.Next
    }
    //fmt.Println(firstPre.Val,secondCur.Val)
    cur=list2
    for cur.Next!=nil{
        cur=cur.Next
    }
    firstPre.Next=list2
    cur.Next=secondCur.Next
    return dummy.Next
}
