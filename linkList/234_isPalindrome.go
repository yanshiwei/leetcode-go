/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head==nil||head.Next==nil{
        return true
    }
    /*
    确定数组列表是否回文很简单，我们可以使用双指针法来比较两端的元素，并向中间移动。一个指针从起点向中间移动，另一个指针从终点向中间移动。故首先转换成数组
    */
    var arr=make([]int,0)
    var cur=head
    for cur!=nil{
        arr=append(arr,cur.Val)
        cur=cur.Next
    }
    for i:=range arr[:len(arr)/2]{
        if arr[i]!=arr[len(arr)-1-i]{
            return false
        }
    }
    return true
}
