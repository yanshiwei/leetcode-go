type MyStack struct {
    arr []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    var st=MyStack{arr:make([]int,0)}
    return st
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.arr=append(this.arr,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    v:=this.arr[len(this.arr)-1]
    this.arr=this.arr[0:len(this.arr)-1]
    return v
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.arr[len(this.arr)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if len(this.arr)<1 {
        return true
    }
    return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
