/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int
    var head,tail *ListNode
    for l1!=nil||l2!=nil{
        var n1 int
        if l1!=nil{
            n1=l1.Val
        }
        var n2 int
        if l2!=nil{
            n2=l2.Val
        }
        sum:=n1+n2+carry
        carry=sum/10
        if head==nil{
            head=&ListNode{
                Val:sum%10,
                Next:nil,
            }
            tail=head
        }else{
            tmp:=&ListNode{
                Val:sum%10,
                Next:nil,
            }
            tail.Next=tmp
            tail=tail.Next
        }
        if l1!=nil{
            l1=l1.Next
        }
        if l2!=nil{
            l2=l2.Next
        }
    }
    if carry>0{
        if head==nil{
            head=&ListNode{
                Val:carry,
                Next:nil,
            }
            tail=head
        }else{
            tmp:=&ListNode{
                Val:carry,
                Next:nil,
            }
            tail.Next=tmp
            tail=tail.Next
        }       
    }
    return head
}
