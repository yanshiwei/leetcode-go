func minimumDeviation(nums []int) int {
    /*
    可以考虑把所有数都化为最大的形式(全部化为所能变成的最大偶数，奇数*2，偶数不变)，然后加入大根堆，维护 mimi 表示当前堆的最小值，堆顶即为最大值。

    为了不断缩小最大值与最小值的差，我们必须不断取出当前堆最大的数，然后将其除以2，重新加入队列。

    这个过程同时维护答案 resres 和 mimi，则有 res = min(res, q.top() - mi)res=min(res,q.top()−mi) ，mi = min(mi, q.top() >> 1)mi=min(mi,q.top()>>1)。

    当堆顶为奇数时，意味着不能再除以2，最大值不可能再改变，差距也就不会再被缩小，此时的 resres 即为答案。
    */
    var res int=1e9
    var curMin int=1e9
    var maxHeap=&IntHeap{}
    heap.Init(maxHeap)
    for i:=range nums{
        if nums[i]%2!=0{
            //奇数
            nums[i]*=2
        }
        heap.Push(maxHeap,nums[i])
        curMin=min(curMin,nums[i])
    }
    for maxHeap.Len()>0{
        curMax:=heap.Pop(maxHeap).(int)
        res=min(res,curMax-curMin)
        if curMax%2!=0{
            break
        }
        curMax/=2
        heap.Push(maxHeap,curMax)
        curMin=min(curMin,curMax)
    }
    return res
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
type IntHeap []int
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){
    if i<len(h)&&j<len(h){
        h[i],h[j]=h[j],h[i]
    }
}
func (h IntHeap)Less(i,j int)bool {
    return h[i]>h[j]//max
}
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(int))
}
func (h IntHeap)Top()int {
    return h[0]
}
func (h *IntHeap)Pop()interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
