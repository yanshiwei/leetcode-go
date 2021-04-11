func busyStudent(startTime []int, endTime []int, queryTime int) int {
    var res int
    for i:=range startTime{
        if startTime[i]<=queryTime&&endTime[i]>=queryTime{
            res++
        }
    }
    return res
}
