type MinStack struct {
    arr []int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    var st =MinStack{arr:make([]int,0)}
    return st
}


func (this *MinStack) Push(x int)  {
    if len(this.arr)<1 {
        this.min=x
    }else {
        if this.min > x {
            this.min=x
        }
    }
    this.arr=append(this.arr,x)
}


func (this *MinStack) Pop()  {
    if len(this.arr)>0 {
        V:=this.arr[len(this.arr)-1]
        if this.min==V {
            this.min=this.arr[0]
            for i:=0;i<len(this.arr)-1;i++ {
                if this.min>this.arr[i] {
                    this.min=this.arr[i]
                }
            }
        }
        this.arr=this.arr[0:len(this.arr)-1]
    }
}


func (this *MinStack) Top() int {
    if len(this.arr)>0 {
        return this.arr[len(this.arr)-1]
    }else {
        return 0
    }
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
