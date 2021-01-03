type MyCircularDeque struct {
    CacheList []int
    Len int
    Head int//队首
    Tail int//正好是队尾下一个
}
/*
1.首先需要理解双端队列的含义：队首和队尾都允许插入元素，也允许删除元素；
2.由于数量可控，采用数组实现最节省空间；
3.队首添加元素是将head做减法，队首删除元素是将head做加法，而队尾类似操作时的下标运算正好相反；
4.为了避免添加删除元素时的数据搬迁损耗，采用循环队列来实现；
5.循环队列需要内部预留一个元素来判断队满，所以队列真实长度是k + 1;
6.虽然名字听着高大上，其实双端队列就是一个数组，所以不论是队首还是队尾，操作元素时唯一复杂的地方就是边界的控制；
    6.1 tail在自增时有可能超出数组下标，所以采用(tail + 1) % len 取余的方法来实现下标的循环使用，head自增时也是如此；
    6.2 head在自减时有可能变成负数，所以采用（head - 1 + len) % len取余的方法来实现下标的循环使用，tail自减时也是如此；

*/
/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
        return MyCircularDeque{
        CacheList:make([]int,k+1),
        Len:k+1,//多加一位标记tail,实际容量只有k,因为this.front == (this.rear+1)%this.cap表示队满,所以需要牺牲一格空间
        Head:0,
        Tail:0,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull(){
        return false
    }
    this.Head=(this.Head-1+this.Len)%this.Len// insert head is minus
    this.CacheList[this.Head]=value
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull(){
        return false
    }
    this.CacheList[this.Tail]=value
    this.Tail=(this.Tail+1)%this.Len
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty(){
        return false
    }
    this.Head=(this.Head+1)%this.Len//delete front is add
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty(){
        return false
    }
    this.Tail=(this.Tail-1+this.Len)%this.Len//delete last is minus
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty(){
        return -1
    }
    return this.CacheList[this.Head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty(){
        return -1
    }
    realTail:=(this.Tail-1+this.Len)%this.Len
    return this.CacheList[realTail]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.Head==this.Tail
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.Head==(this.Tail+1)%this.Len
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
