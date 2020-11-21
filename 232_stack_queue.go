type MyQueue struct {
    arr []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    var qu =MyQueue{arr:make([]int,0)}
    return qu
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.arr=append(this.arr,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    v:=this.arr[0]
    this.arr=this.arr[1:]
    return v
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.arr[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.arr)<1 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
