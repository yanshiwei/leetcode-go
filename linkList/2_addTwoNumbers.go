/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    /*
    由于输入的两个链表都是逆序存储数字的位数的，因此两个链表中同一位置的数字可以直接相加。
    我们同时遍历两个链表，逐位计算它们的和，并与当前位置的进位值相加。具体而言，如果当前两个链表处相应位置的数字为 n1,n2，进位值为carry，则它们的和为n1+n2+carry；其中，答案链表处相应位置的数字为 (n1+n2+carry)mod10，而新的进位值为 (n1+n2+carry)/10
    如果两个链表的长度不同，则可以认为长度短的链表的后面有若干个 0 。
    此外，如果链表遍历结束后，有carry>0，还需要在答案链表的后面附加一个节点，节点的值为carry
    */
    var head,tail *ListNode
    var carry int
    for l1!=nil||l2!=nil{
        var v1 int
        if l1!=nil{
            v1=l1.Val
        }
        var v2 int
        if l2!=nil{
            v2=l2.Val
        }
        var sum=v1+v2+carry
        if head==nil{
            //头节点
            head=&ListNode{
                Val:sum%10,
                Next:nil,
            }
            tail=head
        }else{
            //普通节点
            tail.Next=&ListNode{
                Val:sum%10,
                Next:nil,
            }
            tail=tail.Next
        }
        carry=sum/10
        if l1!=nil{
            l1=l1.Next
        }
        if l2!=nil{
            l2=l2.Next
        }
    }
    if carry>0{
        tail.Next=&ListNode{
            Val:carry,
            Next:nil,
        }
    }
    return head
}
