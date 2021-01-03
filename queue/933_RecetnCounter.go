type RecentCounter struct {
    CacheList []int
    Len int
}


func Constructor() RecentCounter {
    var v RecentCounter
    v.Len=0
    return v
}


func (this *RecentCounter) Ping(t int) int {
    this.CacheList=append(this.CacheList,t)
    var i int
    for i=range this.CacheList{
        if this.CacheList[i]<t-3000{
            //skip
        }else {
            break
        }
    }
    this.Len=len(this.CacheList[i:])
    return this.Len
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
