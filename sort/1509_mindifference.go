func minDifference(nums []int) int {
    /*
    当给定的数组长度不超过4时，我们总可以让所有的数字相同，所以我们直接考虑长度超过 4的数组。
    我们注意到，每次修改必然是将最大值改小，或者将最小值改大，这样才能让最大值与最小值的差尽可能小。
    这样我们只需要找到最大的四个数与最小的四个数即可。当我们删去最小的 k(0≤k≤3) 个数，还需要删去 3-k个最大值。枚举这四种情况即可
    为了保持差值最小，对应有：
    最小值   最大值
    0        m-4
    1        m-3
    2        m-2
    3        m-1
    */
    var m=len(nums)
    if m<=4{
        return 0
    }
    sort.Ints(nums)
    var ret=INTMAX
    for i:=0;i<=3;i++{
        ret=min(ret,nums[m-4+i]-nums[i])
    }
    return ret
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
const INTMAX=int(^(uint(0))>>1)
