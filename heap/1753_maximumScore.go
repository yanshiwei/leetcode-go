func maximumScore(a int, b int, c int) int {
    //为了使得游戏能够尽可能玩下去，我们尽量保证剩余的非空堆多一点，因此，每次都贪心地从石子更多的堆中取石子即可。使用大根堆存储三个堆地石子数量，每次取堆顶的两个（也就是堆中最大地两个值）
    var maxHeap=&IntHeap{}
    heap.Init(maxHeap)
    var res int
    heap.Push(maxHeap,a)
    heap.Push(maxHeap,b)
    heap.Push(maxHeap,c)
    for maxHeap.Len()>=2{
        //取出前2个最大的
        num1:=maxHeap.Top()
        heap.Pop(maxHeap)
        num2:=maxHeap.Top()
        heap.Pop(maxHeap)
        //fmt.Println(num1,num2)
        //减去1
        num1--
        num2--
        res++
        if num1>0{
            heap.Push(maxHeap,num1)
        }
        if num2>0{
            heap.Push(maxHeap,num2)
        }
    }
    return res
}

type IntHeap []int
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool{return h[i]>h[j]}//max
func (h IntHeap)Top()int{return h[0]}
func (h *IntHeap)Push(x interface{}){*h=append(*h,x.(int))}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
