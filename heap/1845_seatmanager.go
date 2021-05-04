type SeatManager struct {
    Setas *IntHeap
    TOPK int
}


func Constructor(n int) SeatManager {
    var m=SeatManager{
        Setas:&IntHeap{},
        TOPK:n,
    }
    heap.Init(m.Setas)
    for i:=1;i<=n;i++{
        heap.Push(m.Setas,i)
    }
    return m
}


func (this *SeatManager) Reserve() int {
    x:=this.Setas.Top()
    heap.Pop(this.Setas)
    return x
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.Setas,seatNumber)
}

type IntHeap []int
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool{return h[i]<h[j]}
func (h IntHeap)Top()int{return h[0]}
func (h*IntHeap)Push(x interface{}){*h=append(*h,x.(int))}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
