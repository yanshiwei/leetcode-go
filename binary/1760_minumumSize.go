func minimumSize(nums []int, maxOperations int) int {
    /*
    二分法，求最值的最值，类似https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    0.maxOperations次操作中单个袋子里的球数目的最大值的最小化
    1.对某个开销值，看它能不能在maxOperations内操作得到。理论上只要Operation足够大，完全可以达到最小开销1。最大开销值越大，需要的Operations越少。最终要得到结果在1~1000_000_000内，连续且有序，所以可以用二分法去找这个符合条件（operations <= maxOperations）的最小的开销
    2.问题的关键是如果判断一个开销值合法，即是不是在小于等于maxOperations内可以操作所有的nums
    因为每次都是重新分给两个新袋子。则对于nums里的每一个数，如果开销为maxCnt，那么就用这个数去除maxCnt，不能整除的话，就是除之后得到的结果如num=10，开销为3，10/3=3,意思是3次可以分成（3，3，3，1）如果是能整除如num=9，开销为3，9/3=3，其实2次就能分完（3，3，3）
    3.类似题目：
    https://leetcode-cn.com/problems/split-array-largest-sum/
    https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    https://leetcode-cn.com/problems/xiao-zhang-shua-ti-ji-hua/
    */
    var left,right int
    left=1//最小是操作后球数目是1
    right=nums[0]//最大是球数目的最大值
    for i:=range nums{
        if right<nums[i]{
            right=nums[i]
        }
    }
    var ans =-1
    for left<=right{
        var mid=left+(right-left)/2
        if canPut(nums,mid,maxOperations){
            //如果单个袋子最大球数目mid个情况下可以在最多maxOperations次完成
            ans=mid
            right=mid-1
        }else{
            left=mid+1
        }
    }
    return ans
}

func canPut(nums[]int,maxCnt,maxOperations int)bool{
    var cnt int//操作次数,maxCnt是单次操作可放入的最多球数
    for i:=range nums{
        cnt+=(nums[i]/maxCnt)
        if nums[i]%maxCnt==0{
            cnt--
        }
        if cnt>maxOperations{
            //未遍历完nums就耗尽里操作次数
            return false
        }
    }
    return true
}
