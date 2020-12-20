func kthSmallest(matrix [][]int, k int) int {
    /*
    var maxHeap []int

    for i:=0;i<len(matrix);i++ {
        for j:=0;j<len(matrix[0]);j++ {
            if len(maxHeap)<k {
                maxHeap=append(maxHeap,matrix[i][j])
                buildHeap(maxHeap)
            }else {
                if matrix[i][j]<maxHeap[0]{
                    maxHeap=maxHeap[1:]
                    maxHeap=append(maxHeap,matrix[i][j])
                    buildHeap(maxHeap)
                }
            }
        }
    }
    */
    //简单的方法超时，改正
    var minHeap []Info
    //将每一行第一个元素入堆，每一行的第一个都是该行的最小值
    for i:=0;i<len(matrix);i++{
        info:=Info{X:i,Y:0,V:matrix[i][0]}
        minHeap=append(minHeap,info)
        buildHeap(minHeap)
    }
    //从堆取K次，取完就是第K小。要找到第k小的数,需要把前k-1个最小的数丢弃掉
    for j:=0;j<k-1;j++ {
        info:=minHeap[0]
        minHeap=minHeap[1:]
        buildHeap(minHeap)
        if info.Y==len(matrix[0])-1 {
            continue
        }else {
            infoNew:=Info{X:info.X,Y:info.Y+1,V:matrix[info.X][info.Y+1]}
            minHeap=append(minHeap,infoNew)
            buildHeap(minHeap)
        }
    }
    return minHeap[0].V
}
func adjustHeap(maxHeap[]Info,start,length int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1{
            break
        }
        if child+1<=length-1&&maxHeap[child].V>maxHeap[child+1].V{
            child++
        }
        if maxHeap[i].V>maxHeap[child].V{
            maxHeap[i],maxHeap[child]=maxHeap[child],maxHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(maxHeap []Info){
    for i:=len(maxHeap)/2-1;i>=0;i-- {
        adjustHeap(maxHeap,i,len(maxHeap))
    }
}

type Info struct {
    X int
    Y int
    V int
}
