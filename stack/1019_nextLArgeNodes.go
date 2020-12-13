/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    var res []int
    if head==nil {
        return res
    }
    var stackIndex,arr []int
    var cur=head
    for cur!=nil {
        arr=append(arr,cur.Val)
        cur=cur.Next
    }
    res=make([]int,len(arr))
    //desc stack to get right bigger
    for i:=range arr {
        for len(stackIndex)>0&&arr[stackIndex[len(stackIndex)-1]]<arr[i] {
            //not  desc
            headIndex:=stackIndex[len(stackIndex)-1]
            res[headIndex]=arr[i]
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
        }
        stackIndex=append(stackIndex,i)
    }
    return res
}
