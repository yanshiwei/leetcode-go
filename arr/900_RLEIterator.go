type RLEIterator struct {
    CacheList []int
    CurIndex int64
}


func Constructor(A []int) RLEIterator {
    var m RLEIterator
    m.CacheList=A
    m.CurIndex=0//偶数标
    return m
}

func (this *RLEIterator) Next(n int) int {
    var m int64=int64(len(this.CacheList))
    for this.CurIndex<m&&this.CacheList[this.CurIndex]<n{
        //n大于当前偶数对应值时，则将当前计数置0
        n-=this.CacheList[this.CurIndex]
        this.CacheList[this.CurIndex]=0
        this.CurIndex+=2//下一个偶数下标
    }
    if this.CurIndex>=m{
        return -1
    }
    this.CacheList[this.CurIndex]-=n
    return this.CacheList[this.CurIndex+1]
}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
