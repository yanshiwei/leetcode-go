func twoSum(numbers []int, target int) []int {
    var sumMap=make(map[int]int)
    var res []int
    for i:=range numbers {
        if v,ok:=sumMap[target-numbers[i]];ok {
            res=append(res,v)
            res=append(res,i+1)
            return res
        }else {
            sumMap[numbers[i]]=i+1
        }
    }
    return res
}
