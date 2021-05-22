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
    var cur=head
    var arr []int
    for cur!=nil{
        arr=append(arr,cur.Val)
        cur=cur.Next
    }
    for i:=0;i<=len(arr)/2;i++{
        if arr[i]!=arr[len(arr)-1-i]{
            //fmt.Println(arr[i],arr[i+len(arr)/2])
            return false
        }
    }
    return true
}
