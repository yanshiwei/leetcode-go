/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head,tail *ListNode
    for l1!=nil&&l2!=nil{
        if l1.Val>l2.Val{
            if head==nil{
                head=&ListNode{
                    Val:l2.Val,
                    Next:nil,
                }
                tail=head
            }else{
                tail.Next=&ListNode{
                    Val:l2.Val,
                    Next:nil,
                }
                tail=tail.Next
            }
            l2=l2.Next
        }else{
            if head==nil{
                head=&ListNode{
                    Val:l1.Val,
                    Next:nil,
                }
                tail=head
            }else{
                tail.Next=&ListNode{
                    Val:l1.Val,
                    Next:nil,
                }
                tail=tail.Next
            }
            l1=l1.Next 
        }
    }
    if l1!=nil{
        if head==nil{
            head=l1
        }else{
            tail.Next=l1
        }
    }
    if l2!=nil{
        if head==nil{
            head=l2
        }else{
            tail.Next=l2
        }
    }
    return head
}
