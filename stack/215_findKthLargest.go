func findKthLargest(nums []int, k int) int {
    //size k min heap
    var heap []int
    for i:=range nums {
        if i<k {
            //push only
            heap=append(heap,nums[i])
            buildHeap(heap)
        }else {
            //pop head
            if nums[i]>heap[0] {
                heap=heap[1:]
                heap=append(heap,nums[i])
                buildHeap(heap)
            }
        }
    }
    return heap[0]
}
func adjustHeap(minHeap []int,start, length int) {
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&minHeap[child]>minHeap[child+1]{
            child++
        }
        if minHeap[i]>minHeap[child] {
            minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
            i=child
        }else {
            break
        }
    }
}
func buildHeap(minHeap []int) {
    for i:=len(minHeap)/2-1;i>=0;i-- {
        adjustHeap(minHeap,i,len(minHeap))
    }
}
