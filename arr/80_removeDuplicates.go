func removeDuplicates(nums []int) int {
    if len(nums)<2 {
        return len(nums)
    }
    var i,j int
    var repeatedCnt =1
    for j=1;j<len(nums);j++ {
        //fmt.Println(nums[j])
        if nums[j]==nums[j-1] {
            repeatedCnt++
            if repeatedCnt>2 {
                continue
            }else {
                i++
                nums[i]=nums[j]
            }
        }else {
            repeatedCnt=1
            i++
            nums[i]=nums[j]
        }
    }
    return i+1
}
