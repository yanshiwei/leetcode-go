func productExceptSelf(nums []int) []int {
    var res=make([]int,len(nums))
    //第i位置之前的乘积
    var k=1
    for i:=range nums {
        res[i]=k
        k*=nums[i]
    }
    k=1
    //第i位置之后的乘积，与之前乘积相乘即可
    for i:=len(nums)-1;i>=0;i-- {
        res[i]*=k//之前乘积res[i]，k是之后乘积
        k*=nums[i]
    }
    return res
}
