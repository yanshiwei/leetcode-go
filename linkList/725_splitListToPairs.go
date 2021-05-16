/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
    /*
    如果链表有 N个结点，则分隔的链表中每个部分中都有 n/k 个结点，且前 N%k 部分有一个额外的结点。
    */
    var res []*ListNode
    var N int
    var cur=root
    for cur!=nil{
        N++
        cur=cur.Next
    }
    segmentNum:=N/k
    remainder:=N%k
    cur=root
    for i:=0;i<k;i++{
        //每一段
        head:=&ListNode{}
        tail:=head
        otherNum:=0
        if i<remainder{
            otherNum=1
        }
        for j:=0;j<segmentNum+otherNum;j++{
            tmp:=&ListNode{Val:cur.Val}
            tail.Next=tmp
            tail=tail.Next
            if cur!=nil{
                cur=cur.Next
            }
        }
        res=append(res,head.Next)
    }
    return res
}
