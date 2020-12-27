func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
    //因为纯利润不会是负数，故原始资本W不会减少
    var maxHeap=&IntHeap{}
    var res int=W
    heap.Init(maxHeap)//store max profit
    var itemSlice []Info
    for i:=range Profits{
        info:=Info{Vprofit:Profits[i],Vcapital:Capital[i]}
        itemSlice=append(itemSlice,info)
    }
    //fmt.Printf("before:%+v\n",itemSlice)
    ////根据Capital从小到大排序，相同Capital的Profits大的排在前面
    //排序的原因，是取得>=W的最大profit 以获得最多的资本去投资更多项目
    sort.Slice(itemSlice,func(i,j int)bool {
        if itemSlice[i].Vcapital!=itemSlice[j].Vcapital {
            return itemSlice[i].Vcapital<itemSlice[j].Vcapital
        }
        return itemSlice[i].Vprofit>itemSlice[j].Vprofit
    })
    //fmt.Printf("after:%+v\n",itemSlice)
    var i int
    for i=0;i<len(itemSlice);i++ {
        if itemSlice[i].Vcapital<=res {
            heap.Push(maxHeap,itemSlice[i].Vprofit)
        }else{
            break
        }
    }

    for maxHeap.Len()>0&&k>0{
        k--
        curMax:=heap.Pop(maxHeap)
        res+=curMax.(int)//当前最大利润，且利润将被添加到你的总资本中。方便后面投资新项目
        //接着上一轮i
        for i<len(itemSlice) {
            if itemSlice[i].Vcapital<=res {
                heap.Push(maxHeap,itemSlice[i].Vprofit)
                i++
            }else{
                break
            }
        }
    }
    return res
}
type Info struct {
    Vprofit int
    Vcapital int
}
type IntHeap []int
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int) {
    h[i],h[j]=h[j],h[i]
}
func (h IntHeap)Less(i,j int)bool {
    return h[i]>h[j]//max heap
}
func (h *IntHeap)Push(x interface{}){
    *h=append(*h,x.(int))
}
func (h *IntHeap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
