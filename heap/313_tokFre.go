func topKFrequent(nums []int, k int) []int {
    var cntMap=make(map[int]int)
    for i:=range nums{
        cntMap[nums[i]]++
    }
    var index int
    var heap []Info//min heap
    var res []int
    for k1,v1:=range cntMap {
        fmt.Println("map",k1,v1) 
        if index<k {
            info:=Info{Data:k1,Cnt:v1}
            heap=append(heap,info)
            buildHeap(heap)
        }else {
            if v1>heap[0].Cnt{
                heap=heap[1:]//pop
                info:=Info{Data:k1,Cnt:v1}
                heap=append(heap,info)
                buildHeap(heap)
            }
        }
        index++
    }
    for i:=range heap {
        res=append(res,heap[i].Data)
    }
    return res
}

type Info struct {
    Data int
    Cnt int
}

func adjustHeap(heap []Info,start,length int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&heap[child].Cnt>heap[child+1].Cnt{
            child++
        }
        if heap[i].Cnt>heap[child].Cnt{
            heap[i],heap[child]=heap[child],heap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(heap []Info){
    for i:=len(heap)/2-1;i>=0;i--{
        adjustHeap(heap,i,len(heap))
    }
}
