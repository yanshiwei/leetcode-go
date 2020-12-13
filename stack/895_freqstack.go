type FreqStack struct {
    ValueFre map[int]int
    FreMapStack map[int][]int
    MaxFre int
}

func Constructor() FreqStack {
    fs:=FreqStack{ValueFre:make(map[int]int),FreMapStack:make(map[int][]int)}
    return fs
}


func (this *FreqStack) Push(x int)  {
    this.ValueFre[x]+=1
    this.MaxFre=max(this.MaxFre,this.ValueFre[x])
    if _,ok:=this.FreMapStack[this.ValueFre[x]];ok==false {
        stack:=make([]int,0)
        stack=append(stack,x)
        this.FreMapStack[this.ValueFre[x]]=stack
    }else {
        this.FreMapStack[this.ValueFre[x]]=append(this.FreMapStack[this.ValueFre[x]],x)
    }
}

func (this *FreqStack) Pop() int {
    x:=this.FreMapStack[this.MaxFre][len(this.FreMapStack[this.MaxFre])-1]//top
    this.ValueFre[x]-=1
    this.FreMapStack[this.MaxFre]=this.FreMapStack[this.MaxFre][0:len(this.FreMapStack[this.MaxFre])-1]//pop
    if len(this.FreMapStack[this.MaxFre])<1 {
        this.MaxFre--
    }
    return x
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
