func shipWithinDays(weights []int, D int) int {
    /*
    值域范围内二分法，船重量范围已知，最小是weights的最大值（只有满足weights中最大才能完成所有），左边界l，
    最大是weights之和，右边界r，从最大最小范围内找到一个，满足D天内搬运，
    在[l,r]二分就可以，计算当运载能力是mid的时候所需要天数day
    若day > D 说明运载能力小了，l = mid+1
    若day <= D,说明运载能力可以，但是不一定符合最小的要求
    直到区间缩小到l == r
    */
    var l=weights[0]
    var r=weights[0]
    var mid int
    var m=len(weights)
    for i:=1;i<m;i++{
        l=max(l,weights[i])
        r+=weights[i]
    }
    var day int
    var todayCurWeight int
    for l<r{
        mid=l+(r-l)/2
        todayCurWeight=0
        day=0
        for i:=0;i<m;i++{
            if todayCurWeight+weights[i]>mid{
                //next day to handle
                day++
                todayCurWeight=weights[i]
            }else{
                //today can handle
                todayCurWeight+=weights[i]
            }
        }
        if day+1<=D{
            //最后不满足todayCurWeight+weights[i]>mid的那天也是一天，故+1 
            r=mid
        }else{
            l=mid+1
        }
    }
    return l
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
