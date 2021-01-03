type MyCircularQueue struct {
    CacheList []int
    Len int
    Head int//队首前一个
    Tail int//正好是队尾
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        CacheList:make([]int,k+1),
        Len:k+1,//多加一位标记tail,实际容量只有k,因为this.front == (this.rear+1)%this.cap表示队满,所以需要牺牲一格空间
        Head:0,
        Tail:0,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull(){
        return false
    }
    this.Tail=(this.Tail+1)%this.Len
    this.CacheList[this.Tail]=value
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty(){
        return false
    }
    this.Head=(this.Head+1)%this.Len
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty(){
        return -1
    }
    return this.CacheList[(this.Head+1)%this.Len]//+1
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty(){
        return -1
    }
    return this.CacheList[this.Tail]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.Head==this.Tail
}


func (this *MyCircularQueue) IsFull() bool {
    return this.Head==(this.Tail+1)%this.Len
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
