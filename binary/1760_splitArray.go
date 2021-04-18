func splitArray(nums []int, m int) int {
    /*
    二分法，求最值的最值，类似https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    0.最大值指的是m个子数组和中的最大值
    1.设m个子数组和中的最大值为10，如果在子数组和中最大值是10的情况下可以产生m个子数组，那么就尝试最大值为9的，直至
    不能产生m个非空子数组。在可产生与不可产生之间的距离就是答案。
    2。问题关键是如何确定当前最大值是否合法，即对num数组在最大值是maxSum下可以分成m和非空连续子数组。
        遍历数组，使用 sum 累加
        如果 sum 超过规定的数值， sum 清零然后重新从当前数字累加，子数组的个数 +1
        如果子数组先用完了，表示失败
        反之跑完了所有数组中的数字，表示成功
    3.类似题目：
    https://leetcode-cn.com/problems/minimum-limit-of-balls-in-a-bag/
    https://leetcode-cn.com/problems/magnetic-force-between-two-balls/solution/liang-qiu-zhi-jian-ci-li-zui-xiao-zhi-zu-g3h3/
    https://leetcode-cn.com/problems/xiao-zhang-shua-ti-ji-hua/
    */
    var left,right int
    //子数组和最小值是数组最大值，最大值是所有数和
    for i:=range nums{
        right+=nums[i]
        if left<=nums[i]{
            left=nums[i]
        }
    }
    var ans=-1
    for left<=right{
        var mid=left+(right-left)/2
        if canSplit(nums,mid,m){
            //该数组和满足m切分
            ans=mid
            right=mid-1
        }else{
            //该数组和不满足切分
            left=mid+1
        }
    }
    return ans
}

func canSplit(nums[]int,maxSum,m int)bool{
    var sum int
    var cnt int
    for i:=range nums{
        sum+=nums[i]
        if sum<=maxSum{
            continue
        }
        sum=nums[i]
        cnt++
        if cnt>=m {
            //子数组先用完了，即不能将数组中所有数字都分割进子数组，即还没分割完，子数组数量就用完了。则表示这个最大值不够大
            return false
        }
    }
    return true
}
