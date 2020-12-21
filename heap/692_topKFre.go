func topKFrequent(words []string, k int) []string {
    var FreMap =make(map[string]int)
    for i:=range words {
        FreMap[words[i]]++
    }
    var minHeap []Info
    for k1,v1:=range FreMap{
        info:=Info{Data:k1,Cnt:v1}
        if len(minHeap)<k {
            minHeap=append(minHeap,info)
            buildHeap(minHeap)
        }else {
            if v1>minHeap[0].Cnt {
                minHeap=minHeap[1:]
                minHeap=append(minHeap,info)
                buildHeap(minHeap)
            }else if v1==minHeap[0].Cnt {
                if compare(k1,minHeap[0].Data)==false {
                    minHeap=minHeap[1:]
                    minHeap=append(minHeap,info)
                    buildHeap(minHeap)
                }
            }
        }
    }
    sort.Sort(StrSlice(minHeap))
    var res []string
    for i:=range minHeap {
        //info:=minHeap[i]
        res=append(res,minHeap[i].Data)
    }
   
    return res
}
type StrSlice []Info
func (s StrSlice)Len()int{return len(s)}
func (s StrSlice)Swap(i,j int){s[i],s[j]=s[j],s[i]}
func (s StrSlice)Less(i,j int)bool {
    if s[i].Cnt>s[j].Cnt {
        return true
    }else if s[i].Cnt<s[j].Cnt  {
        return false
    }
    if compare(s[i].Data,s[j].Data)==false {
        return true
    }
    return false
}
func compare(s1,s2 string)bool {
    length:=min(len(s1),len(s2))
    for i:=0;i<length;i++{
        if s1[i]>s2[i]{
            return true
        }else if s1[i]<s2[i]{
            return false
        }
    }
    if len(s1)>length{
        return true
    }
    return false
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
func adjustHeap(minHeap []Info,start,length int){
    var i=start
    var child int
    for {
        child=2*i+1
        if child>length-1{
            break
        }
        if child+1<=length-1{
            if minHeap[child].Cnt>minHeap[child+1].Cnt {
                child++
            }else if minHeap[child].Cnt==minHeap[child+1].Cnt {
                if compare(minHeap[child].Data,minHeap[child+1].Data)==false {
                    child++
                }
            }
        }
        if minHeap[i].Cnt>minHeap[child].Cnt{
            minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
            i=child
        }else if minHeap[i].Cnt==minHeap[child].Cnt {
            if compare(minHeap[i].Data,minHeap[child].Data)==false {
                minHeap[i],minHeap[child]=minHeap[child],minHeap[i]
                i=child
            }else {
                break
            }
        }else {
            break
        }
    }
}

func buildHeap(minHeap []Info){
    for i:=len(minHeap)/2-1;i>=0;i--{
        adjustHeap(minHeap,i,len(minHeap))
    }
}

type Info struct {
    Data string
    Cnt int
}
