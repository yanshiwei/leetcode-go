func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    //维护一个优先队列，节点内容是到达某个节点用了信息，优先队列按照边权升序排列。

//初始化的时候加入一个起点，每次从有限队列取出一个，然后枚举从这个点开始的所有出边，维护后加入到优先队列
    var edges[][]int
    for i:=0;i<n;i++{
        var one []int
        for j:=0;j<n;j++ {
            one=append(one,-1)
        }
        edges=append(edges,one)
    }
    for i:=range flights {
        oneF:=flights[i]
        edges[oneF[0]][oneF[1]]=oneF[2]
    }
    var minHeap []Info//heap用于存储当前能到达的城市d的相关信息，结构为(cost, d, k)，cost为到达d的花费，k为还能中转的次数
    srcInfo:=Info{ToCityCost:0,CityIdx:src,ResKAfterCity:K+1}//k+1包括自己
    minHeap=append(minHeap,srcInfo)
    
    buildHeap(minHeap)
    for len(minHeap)>0 {
        info:=minHeap[0]
        minHeap=minHeap[1:]
        buildHeap(minHeap)
        if info.CityIdx==dst {
            return info.ToCityCost
        }
        if info.ResKAfterCity>0 {
            srcList:=edges[info.CityIdx]
            for i:=range srcList{
                if srcList[i]>-1 {
                    nextInfo:=Info{ToCityCost:info.ToCityCost+srcList[i],CityIdx:i,ResKAfterCity:info.ResKAfterCity-1}
                    minHeap=append(minHeap,nextInfo)
                    //buildHeap(minHeap)
                }
            }
            buildHeap(minHeap)
        }
    }
    return -1
}
func adjustHeap(minHeap[]Info,start,length int) {
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&minHeap[child].ToCityCost>minHeap[child+1].ToCityCost{
            child++
        }
        if minHeap[i].ToCityCost>minHeap[child].ToCityCost{
            minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(minHeap []Info) {
    for i:=len(minHeap)/2-1;i>=0;i-- {
        adjustHeap(minHeap,i,len(minHeap))
    }
}
type Info struct {
    ToCityCost int
    CityIdx int
    ResKAfterCity int
}
const INTMAX=int(^uint(0)>>1)
