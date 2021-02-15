type MyCalendar struct {
    TimeList [][]int
}

func Constructor() MyCalendar {
    req:=MyCalendar{}
    return req
}


func (this *MyCalendar) Book(start int, end int) bool {
    var m=len(this.TimeList)
    if m<1 {
        var time []int
        time=append(time,[]int{start,end}...)
        this.TimeList=append(this.TimeList,time)
        return true
    }
    for i:=range this.TimeList{
        s,e:=this.TimeList[i][0],this.TimeList[i][1]
        if start<e&&end>s{
            return false
        }
    }
    var time []int
    time=append(time,[]int{start,end}...)
    this.TimeList=append(this.TimeList,time)
    return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
