/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head==nil||head.Next==nil{
        return head
    }
    var dummy=&ListNode{
        Val:0,
        Next:head,
    }
    var pre=dummy
    var cur=head
    var preleftNode,leftNode,rightNode,aftrightNode *ListNode
    var idx int
    for cur!=nil{
        idx++
        if idx==left{
            preleftNode=pre
            leftNode=cur
        }
        if idx==right{
            rightNode=cur
            aftrightNode=rightNode.Next
        }
        pre=cur
        cur=cur.Next
    }
    //fmt.Println(leftNode.Val,rightNode.Val)
    leftNode,rightNode=reverse(leftNode,rightNode)
    preleftNode.Next=leftNode
    rightNode.Next=aftrightNode
    return dummy.Next
}

func reverse(head,tail *ListNode)(*ListNode,*ListNode){
    var pre=tail.Next
    var cur=head
    for pre!=tail{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    return tail,head
}
