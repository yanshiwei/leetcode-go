func kClosest(points [][]int, K int) [][]int {
    var maxHeap[]Info
    for i:=range points {
        point:=points[i]
        info:=Info{X:point[0],Y:point[1]}
        info.Dis=getDis(info)
        if len(maxHeap)<K {
            maxHeap=append(maxHeap,info)
            buildHeap(maxHeap)
        }else {
            if info.Dis<maxHeap[0].Dis {
                maxHeap=maxHeap[1:]
                maxHeap=append(maxHeap,info)
                buildHeap(maxHeap)
            }
        }
    }
    var res[][]int
    for i:=range maxHeap {
        info:=maxHeap[i]
        var point []int
        point=append(point,info.X)
        point=append(point,info.Y)
        res=append(res,point)
    }
    return res
}

func adjustHeap(maxHeap[]Info, start,length int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1{
            break
        }
        if child+1<=length-1&&maxHeap[child].Dis<maxHeap[child+1].Dis {
            child++
        }
        if maxHeap[i].Dis<maxHeap[child].Dis {
            maxHeap[i],maxHeap[child]=maxHeap[child],maxHeap[i]
            i=start
        }else {
            break
        }
    }
}
func buildHeap(maxHeap[]Info){
    for i:=len(maxHeap)/2-1;i>=0;i-- {
        adjustHeap(maxHeap,i,len(maxHeap))
    }
}
type Info struct{
    X int
    Y int
    Dis int
}

func getDis(info Info) int{
    return info.X*info.X+info.Y*info.Y
}
