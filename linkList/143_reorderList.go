/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    //因为链表不支持下标访问，所以我们无法随机访问链表中任意位置的元素。我们利用线性表存储该链表，然后利用线性表可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可
    if head==nil||head.Next==nil{
        return
    }
    var fre=make([]*ListNode,0)
    var cur=head
    for cur!=nil{
        fre=append(fre,cur)
        cur=cur.Next
    }
    var i=0
    var j=len(fre)-1
    for i<j{
        fre[i].Next=fre[j]//0---n
        i++
        if i==j{
            break
        }
        fre[j].Next=fre[i]//n--1
        j--
    }
    fre[j].Next=nil//结束
    return
}
