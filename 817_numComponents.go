/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
    //就是统计给定的列表G里面包含链表中的几段.直接遍历链表，遇到连续的元素算1个块，单个也算一个，遍历统计即可
    var fre=make(map[int]struct{})
    for i:=range nums{
        fre[nums[i]]=struct{}{}
    }
    var cur=head
    var res int
    for cur!=nil{
        if _,ok:=fre[cur.Val];ok==false{
            //not exist
            cur=cur.Next
            continue
        }
        //包含一个时算一块，res+1
        res++
        //如果接着包含连续多个，也算一块res不动
        for cur!=nil{
            if _,ok:=fre[cur.Val];ok==false{
                break
            }
            cur=cur.Next
        }
    }
    return res
}
