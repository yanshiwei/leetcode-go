func minTime(time []int, m int) int {
    /*
    二分法，求最值的最值，类似https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    0.耗时最多指的是m天中做题时间做多的一天对应的耗时
    1.设T=10合法即m天中做题时间最长的一次设10，那么就应该尝试T=9的，直至不能在m天内完成所有的题目。在可完成与不可完成之间的时间就是最佳答案
    2.问题关键是如何确定当前最大值是否合法，即对time数组在最大值是maxTime下可以分成m天内完成所有题目。
    3.类似题目
    https://leetcode-cn.com/problems/minimum-limit-of-balls-in-a-bag/
    https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    https://leetcode-cn.com/problems/split-array-largest-sum/solution/fen-ge-shu-zu-de-zui-da-zhi-zui-xiao-hua-hvv7/
    */
    // 如果天数大于题目 可以0做完 
    if len(time)<=m{
        return 0
    }
    var left,right int
    left=time[0]
    //范围是最小值，所有时间和
    for i:=range time{
        if left>time[i]{
            left=time[i]
        }
        right+=time[i]
    }
    var ans=-1
    for left<=right{
        var mid=left+(right-left)/2
        if canSplit(time,mid,m){
            //如果mid天可以完成所有题目
            ans=mid
            right=mid-1
        }else{
            //如果mid天可以完成所有题目
            left=mid+1
        }
    }
    return ans
}

func canSplit(time []int,maxTime,m int)bool{
    var sum int// 用来记录当前回答问题耗费的时间
    var maxT int// 记录今天的所回答问题的最大值
    var flag=-1// 用来判断是否已经使用过一次求助
    for i:=0;i<len(time);{
        sum+=time[i]
        // 获取今天的最大值
        if time[i]>maxT{
            maxT=time[i]
        }
        // 如果今天耗费的时间大于给定的时间maxTime
        if sum>maxTime{
            // 是否能够使用一次求助 如果可以 就使用一次求助解决耗时最长的一个问题 并把最大值减掉
            // 这样能够回答问题最大化  贪心
            if flag==-1{
                sum-=maxT
                flag=-flag//求助使用完
            }else{
                // 如果已经使用过了
                m--
                if m<=0{
                    // 如果m使用完 但是还没有遍历完成,则m天做不完所有题目
                    return false
                }
                sum=0
                maxT=0
                flag=-flag
                continue
            }
        }
        i++
    }
    return true
}
