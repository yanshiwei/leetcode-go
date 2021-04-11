func minDays(bloomDay []int, m int, k int) int {
    //二分法
    //1<=k<=n
    //已知天数-1<=day<=10^9，按序遍历所有天数 判断是否可以制作花的最少天数（二分法）
    //参考https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/solution/javaban-jiang-jie-si-lu-qing-xi-zhu-shi-jjy72/
    var total=m*k
    var n=len(bloomDay)
    if n<total{
        return -1
    }
    var left=-1
    var right=1000000000
    for left<right{
        var mid=left+(right-left)/2
        if canMake(mid,m,k,bloomDay){
            //如果mid天可以完成那就尝试更小的天数
            right=mid
        }else{
            left=mid+1
        }
    }
    return left
}

func canMake(curDay,m,k int,bloomDay []int)bool{
    var oriK=k
    for i:=range bloomDay{
        oneF:=bloomDay[i]
        if oneF<=curDay{
            ///如元素的值小于daysNum的值说明开花了，可以摘
            k--
            if k==0{
                //一束已经制作好
                m--
                k=oriK//准备下一束
            }
        }else{
            k=oriK//一旦遇到不能摘的立马恢复朵的数量
        }
    }
    if m<=0{
        return true
    }
    return false
}
