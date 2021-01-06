func twoSum(nums []int, target int) []int {
    var mapKey=make(map[int]int)// value idx 
    var res []int
    for i:=range nums {
        if v,ok:=mapKey[target-nums[i]];ok {
            res=append(res,i)
            res=append(res,v)
            return res
        }
        mapKey[nums[i]]=i
    }
    return res
}
