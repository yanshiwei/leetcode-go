func maxSatisfied(customers []int, grumpy []int, X int) int {
    /*
    1.我们可以先将原本就满意的客户加入答案，同时将对应的 customers[i] 变为 0。
    2.之后的问题转化为：在 customers 中找到连续一段长度为 x 的子数组，使得其总和最大。通过窗口法实现
    */
    var m=len(customers)
    var sum int
    for i:=range customers {
        if grumpy[i]==0{
            sum+=customers[i]
            customers[i]=0
        }
    }
    var res,cur int
    //窗口[j,i]
    var j int
    for i:=0;i<m;i++{
        cur+=customers[i]
        if i-j+1>X{
            cur-=customers[j]
            j++
        }
        res=max(res,cur)
    }
    return sum+res
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
