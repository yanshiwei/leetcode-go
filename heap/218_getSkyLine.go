func getSkyline(buildings [][]int) [][]int {
    //核心上高度变化（区分左右顶点）
    //将矩形抽象成两个点，左上角顶点，又上角顶点。将所有顶点按照横坐标进行排序然后遍历
    //遍历时通过最大堆存储当前图像的最高位置，堆顶时所有顶点中的最高的顶点。
    //只要这个点没有被移除堆，则说明这个最高点对应的矩形还没有结束
    //故左顶点加入堆；右顶点，找出堆中相应的左顶点然后移除，意味着整个矩形结束
    //左右通过正负区分
    var res [][]int
    var allPoints=make([][2]int,len(buildings)*2)
    //*2代表矩形的左右两个顶点,内部数组代表X坐标和高度
    var pointIdx int
    for i:=range buildings {
        oneBuild:=buildings[i]
        leftX:=oneBuild[0]
        rightX:=oneBuild[1]
        hight:=oneBuild[2]
        //left point，左顶点存负数
        allPoints[pointIdx][0]=leftX
        allPoints[pointIdx][1]=-hight
        pointIdx++
        //right point，右顶点存正数
        allPoints[pointIdx][0]=rightX
        allPoints[pointIdx][1]=hight
        pointIdx++
    }
    //切片排序
    //根据横坐标对列表排序，相同横坐标的点纵坐标小的排在前面
    sort.Slice(allPoints,func(i,j int)bool{
        if allPoints[i][0]!=allPoints[j][0]{
            //先比较左侧坐标
            return allPoints[i][0]<allPoints[j][0]
        }
        //如果坐标相等 则比较高度
        return allPoints[i][1]<allPoints[j][1]
    })
    maxHeap:=&IntHeap{}
    heap.Init(maxHeap)
    var preHeight int//记录上次keypoint关键点高度
    for i:=range allPoints {
        onePoint:=allPoints[i]
        if onePoint[1]<0 {
            //left push
            heap.Push(maxHeap,-onePoint[1])
        }else {
            //right, find its left
            fmt.Println(maxHeap,onePoint)
            for i:=0;i<maxHeap.Len();i++{
                if maxHeap.Get(i)==onePoint[1] {
                    //find !
                    heap.Remove(maxHeap,i)//not pop
                    break
                }
            }
        }
        //一个矩形完成遍历
        maxV:=maxHeap.Top()//当前矩形最高度
        if maxV!=preHeight {
            // 如果堆的新顶部和上个keypoint高度不一样，上一个更高的矩形已经遍历完成，取下一个次高的，则加入一个新的keypoint
            preHeight=maxV
            res=append(res,[]int{onePoint[0],maxV})
        }//如果相等，则说明PRE还是最高的
    }
    return res
}

type IntHeap []int
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int)bool {
    return h[i]>h[j]//max heap
}
func (h IntHeap)Get(index int)int {
    return h[index]
}
func (h IntHeap)Top()int {
    if len(h)<1 {
        return 0
    }
    return h[0]
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
