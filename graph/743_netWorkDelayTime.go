func networkDelayTime(times [][]int, N int, K int) int {
    var edges[][]int
    //100*100 edge graph
    for i:=0;i<=100;i++ {
        var one []int
        for j:=0;j<=100;j++ {
            one=append(one,-1)
        }
        edges=append(edges,one)
    }
    var dist=make([]int,N+1)//储存K到其他节点最短路径
    for i:=range times {
        one:=times[i]
        edges[one[0]][one[1]]=one[2]
    }
    for i:=range dist{
        dist[i]=INTMAX
    }
    dist[K]=0//第K个节点 节点数从1开始数
    var queue []int
    queue=append(queue,K)
    for len(queue)>0 {
        //下一轮queue非空
        var visited =make(map[int]int)
        for i:=len(queue);i>0;i-- {
            ////可能本轮q为空，但还有节点未被扫到，所以加循环 进入下一轮
            u:=queue[0]
            queue=queue[1:]
            for v:=1;v<=100;v++ {
                if edges[u][v]!=-1&&(dist[u]+edges[u][v])<dist[v] {
                    if _,ok:=visited[v];ok==false {
                        visited[v]++
                        queue=append(queue,v)
                    }
                    dist[v]=dist[u]+edges[u][v]
                }
            }
        }
    }
    var res int=INTMIN
    for i:=1;i<=N;i++{
        res=max(res,dist[i])
    }
    if res==INTMAX {
        return -1
    }
    return res
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
