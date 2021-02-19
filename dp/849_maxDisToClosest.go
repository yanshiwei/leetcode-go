func maxDistToClosest(seats []int) int {
    /*
    left[i]表示空座位i，离坐在i左边的人的最近的距离。如果i位置为空，则
    left[i]=left[i-1]+1;如果i位置坐人啦，则left[i]=0
    right[i]类似
    那么该座位到最近的人的距离为 min(left[i], right[i])
    */
    var m=len(seats)
    var res int
    if m<1 {
        return res
    }
    var left=make([]int,m)
    for i:=range left{
        left[i]=m
    }
    var right=make([]int,m)
    for i:=range right{
        right[i]=m
    }
    for i:=0;i<m;i++{
        if seats[i]==1{
            left[i]=0
        }else if i>0{
            left[i]=left[i-1]+1
        }
    }
    for i:=m-1;i>=0;i--{
        if seats[i]==1{
            right[i]=0
        }else if i<m-1{
            right[i]=right[i+1]+1
        }
    }
    for i:=range seats{
        if seats[i]==0{
            //如果没坐人
            res=max(res,min(left[i],right[i]))
        }
    }
    return res
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
