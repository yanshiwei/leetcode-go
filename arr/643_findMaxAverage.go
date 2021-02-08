func findMaxAverage(nums []int, k int) float64 {
    var minAvg=-10001.0//所给数据范围 [-10,000，10,000]
    if k<1 {
        return minAvg
    }
    if len(nums)<k {
        return minAvg
    }
    for i:=0;i<len(nums);i++{
        end:=i+k
        if end>len(nums){
            break
        }
        tmp:=getAvg(nums[i:end])
        if tmp>minAvg {
            minAvg=tmp
        }
    }
    return minAvg
}
func getAvg(nums[]int)float64{
    var sum int
    for i:=range nums {
        sum+=nums[i]
    }
    return float64(sum)/float64(len(nums))
}
