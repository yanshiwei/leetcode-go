type FrontMiddleBackQueue struct {
    Arr []int
}


func Constructor() FrontMiddleBackQueue {
    var m=FrontMiddleBackQueue{
    }
    return m
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
    this.Arr=append([]int{val},this.Arr...)
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
    pos:=len(this.Arr)/2
    this.Arr=append(this.Arr[:pos],append([]int{val},this.Arr[pos:]...)...)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
    this.Arr=append(this.Arr,val)
}


func (this *FrontMiddleBackQueue) PopFront() int {
    if len(this.Arr)<1{
        return -1
    }
    head:=this.Arr[0]
    this.Arr=this.Arr[1:]
    return head
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
    if len(this.Arr)<1{
        return -1
    }
    pos:=(len(this.Arr)-1)/2
    mid:=this.Arr[pos]
    this.Arr=append(this.Arr[0:pos],this.Arr[pos+1:]...)
    return mid
}


func (this *FrontMiddleBackQueue) PopBack() int {
    if len(this.Arr)<1{
        return -1
    }
    back:=this.Arr[len(this.Arr)-1]
    this.Arr=this.Arr[:len(this.Arr)-1]
    return back
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
