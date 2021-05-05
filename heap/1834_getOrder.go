func getOrder(tasks [][]int) []int {
    /*
    我们需要两个数据结构来实现题目描述中的 CPU 操作。
    1.第一个数据结构负责按照时间顺序将任务分配给 CPU；---->数组
    2.第二个数据结构负责帮助 CPU 在所有任务中选择处理时间最小的那个执行。-->堆
    我们将数组task 中的所有任务按照enqueueTime 升序排序; 按照此顺序消费，
    */
    var idxArr []int
    for i:=0;i<len(tasks);i++{
        idxArr=append(idxArr,i)
    }
    //按照出现的时间排序
    sort.Slice(idxArr,func(i,j int)bool{
        return tasks[idxArr[i]][0]<tasks[idxArr[j]][0]
    })
    var res []int
    var curTime int
    var ptr int//数组遍历的指针
    var minHeap=&Heap{}
    heap.Init(minHeap)
    max:=func(x,y int)int{
        if x<y{
            return y
        }
        return x
    }
    //fmt.Println(idxArr)
    for i:=0;i<len(tasks);i++{
        if minHeap.Len()<1{
            //如果当前没有任务，直接快进到下一个没有分配给 CPU 的那个任务enqueueTime i
            curTime=max(curTime,tasks[idxArr[ptr]][0])
        }
        //将所有不大于当前时间的任务都放入堆
        for ptr<len(tasks)&&tasks[idxArr[ptr]][0]<=curTime{
            info:=Info{Time:tasks[idxArr[ptr]][1],Idx:idxArr[ptr]}
            heap.Push(minHeap,info)
            ptr++
        }
        //fmt.Println(minHeap)
        //选择处理时间最小的任务
        processInfo:=minHeap.Top()
        curTime+=processInfo.Time
        res=append(res,processInfo.Idx)
        heap.Pop(minHeap)
    }
    return res
}

type Heap []Info
func (h Heap)Len()int {return len(h)}
func (h Heap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h Heap)Less(i,j int)bool{
    if h[i].Time!=h[j].Time{
        return h[i].Time<h[j].Time
    }
    return h[i].Idx<h[j].Idx
}
func (h Heap)Top()Info{
    return h[0]
}
func (h*Heap)Push(x interface{}){
    *h=append(*h,x.(Info))
}
func (h*Heap)Pop()interface{}{
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}

type Info struct {
    Time int
    Idx int
}
