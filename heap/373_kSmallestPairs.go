func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    var res[][]int
    var maxHeap []Info
    for i:=range nums1 {
        for j:=range nums2 {
            info:=Info{nums1[i],nums2[j],nums1[i]+nums2[j]}
            if len(maxHeap)<k {
                maxHeap=append(maxHeap,info)
                buidlHeap(maxHeap)
            }else {
                if nums1[i]+nums2[j]<maxHeap[0].Sum {
                    maxHeap=maxHeap[1:]
                    maxHeap=append(maxHeap,info)
                    buidlHeap(maxHeap)
                }
            }
        }
    }
    heapSort(maxHeap)
    for i:=range maxHeap{
        info:=maxHeap[i]
        var tmp []int
        tmp=append(tmp,info.V1)
        tmp=append(tmp,info.V2)
        res=append(res,tmp)
    }
    return res
}

type Info struct {
    V1 int
    V2 int
    Sum int
}

func adjustHeap(maxHeap[]Info,start,length int) {
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&maxHeap[child].Sum<maxHeap[child+1].Sum{
            child++
        }
        if maxHeap[i].Sum<maxHeap[child].Sum{
            maxHeap[i],maxHeap[child]=maxHeap[child],maxHeap[i]
            i=child
        }else {
            break
        }
    }
}

func buidlHeap(maxHeap []Info) {
    for i:=len(maxHeap)/2-1;i>=0;i-- {
        adjustHeap(maxHeap,i,len(maxHeap))
    }
}

func heapSort(maxHeap []Info) {
    buidlHeap(maxHeap)
    for i:=len(maxHeap)-1;i>=0;i-- {
        maxHeap[0],maxHeap[i]=maxHeap[i],maxHeap[0]
        adjustHeap(maxHeap,0,i)
    }
}
