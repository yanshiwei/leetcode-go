func findDisappearedNumbers(nums []int) []int {
    var m=len(nums)
    var res=make([]int,0)
    for i:=range nums {
        idx:=(nums[i]-1)%m//将数组元素对应为索引的位置
        nums[idx]+=m
    }
    //遍历加m后的数组，若数组元素值小于等于n，则说明数组下标值不存在，即消失的数字
    for i:=0;i<m;i++ {
        if nums[i]<=m {
            res=append(res,i+1)
        }
    }
    return res
}
