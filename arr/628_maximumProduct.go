func maximumProduct(nums []int) int {
    if len(nums)<3 {
        return 0
    }
    sort.Ints(nums)
    //若全部正数，则最大三个数相乘
    //若全部负数，则最大三个数相乘
    //若正负都有，可能是最小的两个负数与最大正值乘积
    return max(nums[0]*nums[1]*nums[len(nums)-1],nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
