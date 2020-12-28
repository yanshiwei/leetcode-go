func swimInWater(grid [][]int) int {
    /*
每到一个点，我们再记录所有的周围能走的点，并再次选择所有能走的点中水位最低的点。

0 1 4
2 8 7
3 6 5
此时我们第一步有两个选择 [1,2]，优先选择 1 点，1 点之后有 [4,8] 两点能走，现在所有能走的点为：[2,4,8] （1 点已经走过了）
明显 2 点是水位最低的点，之后又多了 3 点能走（8 点已经记录了）此时所有能走的点为：[3,4,8]（2 点已经走过了）
再走3点，多了6点可以走，此时所有能走的点为[6,7,8]选择第6号点，6号点能达到终点，达成目标
观察我们刚才走过的路线，对应水位最高的点为 6 点，因此答案就是 6 了
    */
    //由起始点开始，每次都走 当前所在的点周围水位最低的点。
    var minHeap=&IntHeap{}
    heap.Init(minHeap)
    var time int //time记录上一次到达地方的时间
    var dx=[]int{0,0,-1,1}
    var dy=[]int{-1,1,0,0}
    first:=Info{X:0,Y:0,Time:grid[0][0]}//左上平台 (0，0) 出发,time=0
    heap.Push(minHeap,first)
    grid[0][0]=-1//-1 表示已经访问了
    for minHeap.Len()>0{
        info:=heap.Pop(minHeap).(Info)
        time=max(time,info.Time)
        if info.X==len(grid)-1&&info.Y==len(grid[0])-1{
            //terminal
            return time
        }
        for i:=0;i<4;i++{
            x:=info.X+dx[i]
            y:=info.Y+dy[i]
            if x<0||y<0||x>=len(grid)||y>=len(grid[0]){
                continue
            }
            if grid[x][y]<0 {
                continue
            }
            other:=Info{X:x,Y:y,Time:grid[x][y]}//左上平台 (0，0) 出发,time=0
            heap.Push(minHeap,other)
            grid[x][y]=-1
        }
    }
    return time
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

type Info struct {
    X int
    Y int
    Time int
}
type IntHeap []Info
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){
    h[i],h[j]=h[j],h[i]
}
func (h IntHeap)Less(i, j int)bool {
    return h[i].Time<h[j].Time//min heap
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
