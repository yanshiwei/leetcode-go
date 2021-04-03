func isNStraightHand(hand []int, W int) bool {
    var m=len(hand)
    if m%W!=0{
        return false
    }
    var fre=make(map[int]int)//元素出现次数
    sort.Ints(hand)
    setCnt:=m/W
    var curSetCnt int
    for i:=range hand{
        fre[hand[i]]++
    }
    //fmt.Println(fre)
    for i:=range hand{
        cnt:=fre[hand[i]]
        if cnt==0{
            // 等于0说明被归到前面集合中了，不用再处理
            continue
        }
        fre[hand[i]]=cnt-1
        //集合的第一个元素是hand[i]，后续需要k-1个连续数
        for j:=1;j<W;j++{
            innerCnt:=fre[hand[i]+j]
            if innerCnt==0{
                //fmt.Println(i,j,hand[i])
                // 等于0就说明以当前num为起点，找不到k个大小的连续集合
                return false
            }
            fre[hand[i]+j]=innerCnt-1
        }
        curSetCnt++
        if curSetCnt==setCnt{
            return true
        }
    }
    return true
}
