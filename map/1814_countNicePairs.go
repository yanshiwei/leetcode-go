func countNicePairs(nums []int) int {
    /*
    由公式 nums[i] + rev(nums[j]) == nums[j] + rev(nums[i]) 可以得出
    nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
    先求出每个元素和自己的反转的差, 存入哈希表中, 最后求出每个差值的对子数, 相加即可
    */
    rev:=func(x int)int {
        var res int
        for x>0{
            mod:=x%10
            res=res*10+mod
            x/=10
        }
        return res
    }
    var mod=1000000000+7
    var res int
    var fre=make(map[int]int)
    for i:=range nums{
        sum:=nums[i]-rev(nums[i])
        fre[sum]++
    }
    //每个差值的对子数
    //C(n,2)=n*(n-1)/2
    //如差值为-3的频次为3，则3个元素之间两两组和，一共C(3，2)
    for _,cnt:=range fre{
        res+=(cnt*(cnt-1)/2)
    }
    //fmt.Println(fre)
    return res%mod
}
