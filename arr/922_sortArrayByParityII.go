func sortArrayByParityII(nums []int) []int {
    var res =make([]int,len(nums))
    var oddIndex int=1//奇数
    var evenIndex int//偶数
    //sort.Ints(nums)
    for i:=range nums {
        if nums[i]%2==0{
            res[evenIndex]=nums[i]
            evenIndex+=2
        }else{
            res[oddIndex]=nums[i]
            oddIndex+=2
        }
    }
    return res
}
