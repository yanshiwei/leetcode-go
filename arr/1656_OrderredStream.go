type OrderedStream struct {
    Arr []string
    Ptr int
}


func Constructor(n int) OrderedStream {
    var m=OrderedStream{
        Arr:make([]string,n+1),
        Ptr:1,
    }
    return m
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    var res []string
    this.Arr[idKey]=value
    if this.Ptr!=idKey{
        return res
    }
    //fmt.Println(idKey,this.Ptr,this.Arr)
    for i:=idKey;i<len(this.Arr);i++{
        if this.Arr[i]==""{
            break
        }
        res=append(res,this.Arr[i])
        this.Ptr+=1
    }
    return res
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
