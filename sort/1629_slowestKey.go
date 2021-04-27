func slowestKey(releaseTimes []int, keysPressed string) byte {
    if len(releaseTimes)==1{
        return keysPressed[0]
    }
    var fre=make([]byte,0)//loc-cnt
    var costTime=releaseTimes[0]
    //找出最长时间
    for i:=1;i<len(releaseTimes);i++{
        delta:=releaseTimes[i]-releaseTimes[i-1]
        if delta>costTime{
            costTime=delta
        }
    }
    for i:=1;i<len(releaseTimes);i++{
        delta:=releaseTimes[i]-releaseTimes[i-1]
        if delta==costTime{
            fre=append(fre,keysPressed[i])
        }
    }
    //最大值就是第一个
    if costTime==releaseTimes[0]{
        fre=append(fre,keysPressed[0])
    }
    sort.Slice(fre,func(i,j int)bool{
        return fre[i]>fre[j]
    })
    //fmt.Println(fre,costTime)
    return fre[0]
}
