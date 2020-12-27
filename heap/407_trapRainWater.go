func trapRainWater(heightMap [][]int) int {
    /*
    创建一个优先级队列，该队列中是将高度小的放在队首。先将四周一圈的格子加入队列中，并将这些格子的状态设为已访问过。下面开始BFS，从队列的头部取出元素，并用其高度更新海平面level，然后从该格子向四周寻找格子，如果某格子未被访问过，那么将其状态设为已访问过，然后判断如果其高度小于海平面，那么把它们的高度差加到ans中，否则不操作。接下来将该格子加入到优先级队列中。如此进行，当优先级队列为空时退出循环，返回ans
    */
    if len(heightMap)<1||len(heightMap[0])<1 {
        return 0
    }
    var minHeap =&IntHeap{}
    heap.Init(minHeap)
    var visited=make([][]bool,len(heightMap))
    for i:=range visited {
        visited[i]=make([]bool,len(heightMap[0]))
    }
    //四周肯定装不了，先将四周入堆
    for i:=0;i<len(heightMap);i++ {
        for j:=0;j<len(heightMap[0]);j++ {
            if i==0||j==0||i==len(heightMap)-1||j==len(heightMap[0])-1 {
                info:=Info{X:i,Y:j,V:heightMap[i][j]}
                visited[i][j]=true
                heap.Push(minHeap,info)
            }
        }
    }
    //bfs
    var level int=INTMIN
    var dx=[]int{-1,0,1,0}//左 下 右 上
    var dy=[]int{0,1,0,-1}
    var res int
    for minHeap.Len()>1{
        curMin:=heap.Pop(minHeap).(Info)
        //更新海平面
        level=max(level,curMin.V)
        //bfs探索该最小点上下左右四个方向周围没访问过的新节点入堆
        for i:=0;i<4;i++ {
            x:=curMin.X+dx[i]
            y:=curMin.Y+dy[i]
            if x<0||x>=len(heightMap)||y<0||y>=len(heightMap[0]){
                continue
            }
            if visited[x][y] {
                continue
            }
            if heightMap[x][y]<level {
                res+=(level-heightMap[x][y])
            }
            visited[x][y]=true
            info:=Info{X:x,Y:y,V:heightMap[x][y]}
            heap.Push(minHeap,info)
        }
    }
    return res
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
type Info struct {
    X int
    Y int
    V int
}
type IntHeap []Info
func (h IntHeap)Len()int{return len(h)}
func (h IntHeap)Swap(i,j int){
    if i<len(h)&&j<len(h){
        h[i],h[j]=h[j],h[i]
    }
}
func (h IntHeap)Less(i,j int)bool {
    if i<len(h)&&j<len(h){
        return h[i].V<h[j].V//min
    }
    return true
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
