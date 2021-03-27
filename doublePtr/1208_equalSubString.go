func equalSubstring(s string, t string, maxCost int) int {
    /*
    对于每一对下标相等的字符，s[i]和t[i] 求出对应的距离
    接着问题就转化为：在一个数组中，在连续子数组的和小于等于 maxCost 的情况下，找到最长的连续子数组长度。
    因此可以用滑动窗口解题
    */
    var dis=make([]int,len(s))//index-dis
    var minDis int =27
    for i:=range s{
        oneDis:=abs(s[i],t[i])
        dis[i]=oneDis
        if minDis>oneDis{
            minDis=oneDis
        }
    }
    //fmt.Println(minDis,dis)
    var start, end int
    var windowSum int
    var res int
    for end<len(s){
        windowSum+=dis[end]
        for windowSum>maxCost{
            res=max(res,end-start)
            windowSum-=dis[start]
            start++
            //通过start减少
        }
        end++
        //通过end增加
    }
    //fmt.Println(res,start,end,windowSum)
    res=max(res,end-start)
    return res
}

func abs(a,b byte)int{
    if a>b{
        return int(a-b)
    }
    return int(b-a)
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
