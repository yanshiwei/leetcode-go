/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNodeHeap[]*ListNode
 func (h ListNodeHeap)Len()int{return len(h)}
 func (h ListNodeHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
 func (h ListNodeHeap)Less(i,j int)bool {
     return h[i].Val<h[j].Val//min heap
 }
 func (h *ListNodeHeap)Push(x interface{}){
     *h=append(*h,x.(*ListNode))
 }
 func (h *ListNodeHeap)Pop()interface{}{
     old:=*h
     n:=len(old)
     x:=old[n-1]
     *h=old[0:n-1]//inner Pop has swap 0 and n-1
     return x
 }
func mergeKLists(lists []*ListNode) *ListNode {
    minHeap:=ListNodeHeap{}
    heap.Init(&minHeap)
    for i:=range lists{
        oneList:=lists[i]
        if oneList!=nil{
            v:=oneList
            heap.Push(&minHeap,v)//每次push都排序
            //插入每个链表的头节点
        }
    }
    var res =&ListNode{}
    var cur=res
    for minHeap.Len()>0{
        curMin:=heap.Pop(&minHeap).(*ListNode)//每次pop都排序
        cur.Next=curMin
        cur=curMin
        if curMin.Next!=nil{
            heap.Push(&minHeap,curMin.Next)
        }
    }
    return res.Next
}
