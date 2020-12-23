func reorganizeString(S string) string {
    /*
    每次从优先队列中取队首的两个映射对儿处理，因为要拆开相同的字母，这两个映射对儿肯定是不同的字母，可以将其放在一起，之后需要将两个映射对儿中的次数自减1，如果还有多余的字母，即减1后的次数仍大于0的话，将其再放回最大堆。由于是两个两个取的，所以最后循环退出后，有可能优先队列中还剩下了一个映射对儿，此时将其加入结果 res 即可。而且这个多余的映射对儿一定只有一个字母了，因为提前判断过各个字母的出现次数是否小于等于总长度的一半，按这种机制来取字母，不可能会剩下多余一个的相同的字母
    */
    var freMap=make(map[byte]int)
    var maxHeap []Info
    for i:=range S {
        freMap[S[i]]++
    }
    for k1,v1:=range freMap {
        if v1>(len(S)+1)/2{
            return ""
            //某个字母出现的频率大于总长度的一半了，那么必然会有两个相邻的字母出现
        }
        info:=Info{Data:k1,Cnt:v1}
        maxHeap=append(maxHeap,info)
        buildHeap(maxHeap)
    }
    var res string
    for len(maxHeap)>=2{
        first:=maxHeap[0]
        second:=maxHeap[1]
        res+=string(first.Data)
        res+=string(second.Data)
        //两个数字肯定不同
        maxHeap=maxHeap[2:]
        buildHeap(maxHeap)
        first.Cnt--
        if first.Cnt>0 {
            maxHeap=append(maxHeap,first)
            buildHeap(maxHeap)
        }
        second.Cnt--
        if second.Cnt>0 {
            maxHeap=append(maxHeap,second)
            buildHeap(maxHeap) 
        }
    }
    if len(maxHeap)>0 {
        res+=string(maxHeap[0].Data)
    }
    return res
}
func adjustHeap(heap []Info,start ,length int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1 {
            break
        }
        if child+1<=length-1&&heap[child].Cnt<heap[child+1].Cnt{
            child++
        }
        if heap[i].Cnt<heap[child].Cnt{
            heap[i],heap[child]=heap[child],heap[i]
            i=child
        }else {
            break
        }
    }
}

func buildHeap(heap []Info){
    for i:=len(heap)/2-1;i>=0;i-- {
        adjustHeap(heap,i,len(heap))
    }
}
type Info struct {
    Data byte
    Cnt int
}
