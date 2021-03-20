type SnapshotArray struct {
    Fre map[int]map[int]int
    SnapIdx int
    Arr []int
}


func Constructor(length int) SnapshotArray {
    var m =SnapshotArray{
        Fre:make(map[int]map[int]int),
        SnapIdx:0,
        Arr:make([]int,length),
    }
    return m
}


func (this *SnapshotArray) Set(index int, val int)  {
    if _,ok:=this.Fre[this.SnapIdx];ok==false{
        this.Fre[this.SnapIdx]=make(map[int]int)
    }
    this.Fre[this.SnapIdx][index]=val
    this.Arr[index]=val
}


func (this *SnapshotArray) Snap() int {
    old:=this.SnapIdx
    this.SnapIdx++
    this.Fre[this.SnapIdx]=make(map[int]int)//拍一次快照就是创建一个新的字典元素
    return old
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    //fmt.Println(this.Fre)
    v,ok:=this.Fre[snap_id][index]
    for ok==false{
        if snap_id==0{
            return 0
        }
        snap_id--
        v,ok=this.Fre[snap_id][index]
    }
    return v
}
const INT_MAX=int(^uint(0)>>1)
const INT_MIN=^INT_MAX

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
