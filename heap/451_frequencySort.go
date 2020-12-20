func frequencySort(s string) string {
    var cntMap=make(map[rune]int)
    for i:=range s {
        cntMap[rune(s[i])]++
    }
    var minHeap []Info
    for k,v:=range cntMap {
        info:=Info{Letter:k,Cnt:v}
        minHeap=append(minHeap,info)
    }
    heapSort(minHeap)
    var res string
    for i:=range minHeap {
        info:=minHeap[i]
        for j:=0;j<info.Cnt;j++{
            res+=string(info.Letter)
        }
    }
    return res
}
type Info struct {
    Letter rune
    Cnt int
}

func adjustHeap(minHeap[]Info,start,length int) {
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1{
            break
        }
        if child+1<=length-1&&minHeap[child].Cnt>minHeap[child+1].Cnt{
            child++
        }
        if minHeap[i].Cnt>minHeap[child].Cnt{
            minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buidlHeap(minHeap[]Info){
    for i:=len(minHeap)/2-1;i>=0;i--{
        adjustHeap(minHeap,i,len(minHeap))
    }
}

func heapSort(minHeap []Info){
    buidlHeap(minHeap)
    for i:=len(minHeap)-1;i>=0;i--{
        minHeap[0],minHeap[i]=minHeap[i],minHeap[0]
        adjustHeap(minHeap,0,i)
    }
}
