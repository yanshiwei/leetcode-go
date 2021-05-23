/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var smallDummy=&ListNode{
        Val:0,
        Next:nil,
    }
    var small=smallDummy
    var bigDummy=&ListNode{
        Val:0,
        Next:nil,
    }
    var big=bigDummy
    var cur=head
    for cur!=nil{
        if cur.Val<x{
            small.Next=cur
            small=small.Next
        }else{
            big.Next=cur
            big=big.Next
        }
        cur=cur.Next
    }
    //合并
    big.Next=nil
    small.Next=bigDummy.Next
    return smallDummy.Next
}
