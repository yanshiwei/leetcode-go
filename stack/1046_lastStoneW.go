func lastStoneWeight(stones []int) int {
    if len(stones)<1 {
        return 0
    }
    if len(stones)<2 {
        return stones[0]
    }
    buildHeap(stones)
    for len(stones)>1{
        maxV:=stones[0]
        stones=stones[1:]
        buildHeap(stones)
        secondV:=stones[0]
        stones=stones[1:]
        buildHeap(stones)
        if maxV!=secondV {
            delat:=abs(maxV-secondV)
            stones=append(stones,delat)
            buildHeap(stones)
        }
    }
    if len(stones)<1 {
        return 0
    }
    return stones[0]
}

func max(x,y int)(int,int) {
    if x<y {
        return x,y
    }
    return y,x
}

func abs(x int) int{
    if x<0 {
        return -x
    }
    return x
}

func adjustHeap(maxHeap []int,start ,lenght int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>lenght-1 {
            break
        }
        if child+1<=lenght-1&&maxHeap[child]<maxHeap[child+1] {
            child++
        }
        if maxHeap[i]<maxHeap[child] {
            maxHeap[i],maxHeap[child]=maxHeap[child],maxHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(maxHeap []int) {
    for i:=len(maxHeap)/2-1;i>=0;i-- {
        adjustHeap(maxHeap,i,len(maxHeap))
    }
}
