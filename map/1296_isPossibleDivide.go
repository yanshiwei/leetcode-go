func isPossibleDivide(nums []int, k int) bool {
    var m=len(nums)
    if m%k!=0{
        return false
    }
    var fre=make(map[int]int)//元素出现次数
    sort.Ints(nums)
    setCnt:=m/k
    var curSetCnt int
    for i:=range nums{
        fre[nums[i]]++
    }
    //fmt.Println(fre)
    for i:=range nums{
        cnt:=fre[nums[i]]
        if cnt==0{
            // 等于0说明被归到前面集合中了，不用再处理
            continue
        }
        fre[nums[i]]=cnt-1
        //集合的第一个元素是nums[i]，后续需要k-1个连续数
        for j:=1;j<k;j++{
            innerCnt:=fre[nums[i]+j]
            if innerCnt==0{
                //fmt.Println(i,j,nums[i])
                // 等于0就说明以当前num为起点，找不到k个大小的连续集合
                return false
            }
            fre[nums[i]+j]=innerCnt-1
        }
        curSetCnt++
        if curSetCnt==setCnt{
            return true
        }
    }
    return true
}
