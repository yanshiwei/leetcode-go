func maximumSwap(num int) int {
    if num==0 {
        return num
    }
    str:=strconv.Itoa(num)
    var strSlice,oldSlice SliceByte
    for i:=range str {
        strSlice=append(strSlice,byte(str[i]))
        oldSlice=append(oldSlice,byte(str[i]))
    }
    sort.Sort(strSlice)
    var i1,i2,i3 int
    for i:=0;i<len(strSlice);i++ {
        if oldSlice[i]!=strSlice[i]{
            i1=i
            i2=i
            break
        }
    }
      //从原始字符串找到不相等字符对应的最后一个位置,r如果不是最后一个位置则不是最大
    for i:=0;i<len(oldSlice);i++ {
        if oldSlice[i]==strSlice[i2]{
            i3=i
        }
    }
    t1:=oldSlice[i1]
    oldSlice[i1]=strSlice[i2]
    oldSlice[i3]=t1
    var res int
    for i:=range oldSlice{
        res=int(oldSlice[i]-'0')+res*10
    }
    return res
}

type SliceByte []byte
func (h SliceByte)Len()int {
    return len(h)
}

func (h SliceByte)Less(i,j int)bool {
    return h[i]>h[j]
}

func (h SliceByte)Swap(i,j int) {
    h[i],h[j]=h[j],h[i]
}
