type StockSpanner struct {
    st []Info
    curIndex int
}

type Info struct {
    value int
    count int
}

func Constructor() StockSpanner {
    var res =StockSpanner{st:make([]Info,0)}
    res.curIndex=0
    return res
}


func (this *StockSpanner) Next(price int) int {
    var res int=1
    if this.curIndex==0 {
        //first
        info:=Info{value:price,count:1}
        this.st=append(this.st,info)//push to store
        this.curIndex++
        return res
    }else{
        var sum int=1
        //bigger so that beofore is reset
        if price>=this.st[len(this.st)-1].value {
            for len(this.st)>0&&price>=this.st[len(this.st)-1].value {
                head:=this.st[len(this.st)-1]
                sum+=head.count
                this.st=this.st[0:len(this.st)-1]//pop
            }
            //pop many ele to stay reset
            info:=Info{value:price,count:sum}
            this.st=append(this.st,info)//push to store
        }else {
            info:=Info{value:price,count:sum}
            this.st=append(this.st,info)//push to store
        }
        res=sum
    }
    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
