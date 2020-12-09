type CustomStack struct {
    St []int
    Cap int
}


func Constructor(maxSize int) CustomStack {
    obj:=CustomStack{St:make([]int,0),Cap:maxSize}
    return obj
}


func (this *CustomStack) Push(x int)  {
    length:=len(this.St)
    if length<this.Cap {
        this.St=append(this.St,x)
    }
}


func (this *CustomStack) Pop() int {
    if len(this.St)<1 {
        return -1
    }
    v:=this.St[len(this.St)-1]
    this.St=this.St[0:len(this.St)-1]//pop
    return v
}


func (this *CustomStack) Increment(k int, val int)  {
    for i:=0;i<len(this.St);i++ {
        if i<k {
            this.St[i]=this.St[i]+val
        }
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
