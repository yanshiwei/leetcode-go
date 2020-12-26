func maxSlidingWindow(nums []int, k int) []int {
    var maxHeap=&IntHeap{}
    heap.Init(maxHeap)
    var res []int
    for i:=range nums{
       
        info:=Info{Data:nums[i],Idx:i}
         /*
        //窗口K 每次右边移动一步就删一个
        if maxHeap.Len()<k {
            heap.Push(maxHeap,info)
            if maxHeap.Len()==k {
                curMax:=maxHeap.Top()
                res=append(res,curMax.Data)
                //deleta before i-(k-1)
                heap.Remove(maxHeap,i-(k-1))
            }
        }
        */
        //time out above 
        heap.Push(maxHeap,info)
        if i>=k-1 {
            //移除当前'过期'的最大值即可，不用窗口K 每次右边移动一步就删一个
            removeIdx:=i-k
            for maxHeap.Len()>0&&maxHeap.Top().Idx<=removeIdx {
                //loop 保证删除干净
                heap.Pop(maxHeap)
            }
            res=append(res,maxHeap.Top().Data)
        }
    }
    return res
}
type Info struct {
    Data int
    Idx int
}
type IntHeap []Info
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){
    if i<len(h)&&j<len(h) {
        h[i],h[j]=h[j],h[i]
    }
}
func (h IntHeap)Less(i,j int)bool {
    if i<len(h)&&j<len(h){
        return h[i].Data>h[j].Data//max heap
    }
    return false
}
func (h IntHeap)Top()Info{
    return h[0]
}
func (h IntHeap)Get(idx int)Info {
    return h[idx]
}
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(Info))
}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
