type KthLargest struct {
    MinHeap []int
    Cap int
}

func adjustHeap(heap []int, start int, length int) {
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&heap[child]>heap[child+1] {
            child+=1
        }
        if heap[i]>heap[child] {
            heap[i],heap[child]=heap[child],heap[i]
            i=child//child update
        }else {
            break
        }
    }
}
func buildHeap(heap []int) {
    //half before
    for i:=len(heap)/2-1;i>=0;i-- {
        adjustHeap(heap,i,len(heap))
    }
}
func Constructor(k int, nums []int) KthLargest {
    kt:=KthLargest{MinHeap:make([]int,0),Cap:k}
    for i:=range nums {
        kt.Add(nums[i])
    }
    return kt
}


func (this *KthLargest) Add(val int) int {
    if len(this.MinHeap)<this.Cap {
        //not full push 到尾部
        this.MinHeap=append(this.MinHeap,val)
        buildHeap(this.MinHeap)
    }else {
        //fmt.Println(this.MinHeap)
        if val>this.MinHeap[0] {
            //弹出堆顶元素
            this.MinHeap=this.MinHeap[1:]
            this.MinHeap=append(this.MinHeap,val)
            //update
            buildHeap(this.MinHeap)
        }
    }
    return this.MinHeap[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
