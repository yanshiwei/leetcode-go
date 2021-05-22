/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    /*
    1.A长度为 a, B长度为b， 假设存在交叉点，此时 A到交叉点距离为 c， 而B到交叉点距离为d
    2.后续交叉后长度是一样的，那么就是 a-c = b-d -> a+d = b+c
    3.这里意味着只要分别让A和B额外多走一遍B和A，那么必然会走到交叉，注意这里边缘情况是，大家都走到null依然没交叉直接返回null
    */
    var curA=headA
    var curB=headB
    for curA!=curB{
        if curA!=nil{
            //curA还没有走完A
            curA=curA.Next
        }else{
            //curA已经走完了A，开始走B
            curA=headB
        }
        if curB!=nil{
            //curB还没有走完B
            curB=curB.Next
        }else{
            //curB已经走完了B，开始走a
            curB=headA
        }
    }
    return curA
}
