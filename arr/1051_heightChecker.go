func heightChecker(heights []int) int {
    var sortArr=make([]int,len(heights))
    for i:=range heights{
        sortArr[i]=heights[i]
    }
    sort.Ints(sortArr)
    var res int
    for i:=range heights{
        if heights[i]!=sortArr[i]{
            res++
        }
    }
    return res
}
