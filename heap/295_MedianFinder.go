type MedianFinder struct {
    minHeap *IntHeap
    maxHeap *IntHeap
}

/*
用一个最大堆存放比中位数小（或等于）的元素，用一个最小堆存放比中位数大（或等于）的元素。这里关键的方法是插入，每当要插入一个元素时，根据判断条件将它插入最大堆或是最小堆，并更新最大堆和最小堆（更新的逻辑很重要），使得最大堆和最小堆中元素的个数之差不超过1，这样中位数就是最大堆或最小堆的堆顶元素。具体而言，当最大堆和最小堆中元素个数不同（个数相差为1）时，元素个数多的那个堆的堆顶元素即为中位数；如果两者元素个数相同，那么中位数可以是最大堆和最小堆的堆顶元素的值取平均。

插入流程（以大堆为基准）
*/
/** initialize your data structure here. */
func Constructor() MedianFinder {
    minH:=&IntHeap{mode:true}
    heap.Init(minH)
    maxH:=&IntHeap{mode:false}
    heap.Init(maxH)
    var heapVal =MedianFinder{minHeap:minH,maxHeap:maxH}
    return heapVal
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxHeap.Len()<1 {
        heap.Push(this.maxHeap,num)
        return
    }
    if this.maxHeap.Len()==this.minHeap.Len(){
        //1.big_heap.size() == small_heap.size()
        //比较最大堆的堆顶元素与将放入的元素大小，如果小于其堆顶的元素，则直接放入到最大堆中，否则放入到最小堆中
        if num<this.maxHeap.Top(){
            //Num < MaxHeapTop
            //比大堆顶小
            heap.Push(this.maxHeap,num)
        }else {
            heap.Push(this.minHeap,num)
        }
    }else if this.maxHeap.Len()>this.minHeap.Len() {
        //2.big_heap.size() > small_heap.size()
        //将放入的元素与最大堆的堆顶元素进行比较，如果大于堆顶元素，直接放入最小堆，否则，将最大堆的堆顶元素push到最小堆中，再将其弹出最大堆，最后将要放入的元素放入到最大堆中
        if num>this.maxHeap.Top() {
            //大于大顶堆元素，直接到小顶堆  
            heap.Push(this.minHeap,num)
        }else {
            //大堆长度=小堆长+1
            heap.Push(this.minHeap,heap.Pop(this.maxHeap))
            heap.Push(this.maxHeap,num)
        }
    }else {
        //3.big_heap.size() < small_heap.size()
        //比较最小堆的堆顶元素与放入元素的大小，如果小于最小堆的堆顶元素，则直接放入到最大堆中，否则，将最小堆的堆顶元素放入到最大堆中，再将堆顶元素弹出，将要放入的元素放入到最小堆
        if num<this.minHeap.Top() {
            heap.Push(this.maxHeap,num)
        }else {
            //小堆长度=大堆长+1
            heap.Push(this.maxHeap,heap.Pop(this.minHeap))
            heap.Push(this.minHeap,num)
        }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len()<this.minHeap.Len() {
        return float64(this.minHeap.Top())
    }else if this.maxHeap.Len()>this.minHeap.Len(){
        return float64(this.maxHeap.Top())
    }else {
        return (float64(this.minHeap.Top())+float64(this.maxHeap.Top()))/2
    }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
 type IntHeap struct{
     heap []int
     mode bool////true是小根堆，false是大根堆
 }
 func (h IntHeap)Len()int{return len(h.heap)}
 func (h IntHeap)Swap(i,j int){
     if i<len(h.heap)&&j<len(h.heap){
         h.heap[i],h.heap[j]=h.heap[j],h.heap[i]
     }
 }
 func (h IntHeap)Less(i,j int)bool {
     if i<len(h.heap)&&j<len(h.heap) {
         if h.mode {
             return h.heap[i]<h.heap[j]//min
         }else {
             return h.heap[i]>h.heap[j]//min
         }
     }else {
         return true
     }
 }
 func (h IntHeap)Top()int {
     if h.Len()>0 {
         return h.heap[0]
     }
     return 0
 }
 func (h *IntHeap)Push(x interface{}) {
     h.heap=append(h.heap,x.(int))
 }
 func (h *IntHeap)Pop()interface{}{
     old:=h
     n:=len(old.heap)
     x:=old.heap[n-1]
     h.heap=old.heap[0:n-1]
     return x
 }
