func minRefuelStops(target int, startFuel int, stations [][]int) int {
    //最少加油次数
    //我们不用去考虑还剩多少油，我们直接考虑，总的油还可以走多远
    //和爬梯子问题一样，延迟处理，先假设不加油可以总很远。走到一个加油站，先不用加油，把他存到堆里。用的时候再拿出来
    //当我们无法到达最近的加油站时，我们就需要加油了，为了保证加油次数最少，首先拿出最多的油去加，故最大顶堆
    if startFuel>=target {
        return 0
    }
    var maxHeap=&IntHeap{}// 把所以没有加的油存到堆中
    heap.Init(maxHeap)
    var addFuelCnt int
    var curFuel=startFuel
    //延迟处理所有加油站
    for i:=range stations{
        for curFuel<stations[i][0]{
            //无法到达最近的加油站时 需要加油
            //这里看你不能用if 因为可能下一个加油站比较远，需要加多次才能到达
            if maxHeap.Len()<1 {
                //无油可以加
                return -1
            }
            curFuel+=heap.Pop(maxHeap).(int)//拿出最多的油去加
            addFuelCnt++
        }
        heap.Push(maxHeap,stations[i][1])//加入
    }
    //消费存储的油
    for curFuel<target {
        if maxHeap.Len()<1 {
            return -1
        }
        curFuel+=heap.Pop(maxHeap).(int)//拿出最多的油去加
        addFuelCnt++
    }
    return addFuelCnt
}
type IntHeap []int
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
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
    old:=*h
    n:=len(old)
    x:=old[n-1]
    *h=old[0:n-1]
    return x
}
