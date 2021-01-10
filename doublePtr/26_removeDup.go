func removeDuplicates(nums []int) int {
    /*
    放置两个指针 i 和 j，j更快一点
    如果a[i]=a[j]则更新j
    直到a[i]!=a[j]则更新i，重复元素已经跳过，
    */
    var res int
    if len(nums)<1{
        return res
    }
    var i int
    var j int
    for j=1;j<len(nums);j++{
        if nums[i]==nums[j]{
            continue
        }else {
            i++
            //新数组下标
            nums[i]=nums[j]
        }
    }
    return i+1
}
