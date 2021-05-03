func findTheWinner(n int, k int) int {
    var members=make([]int,n)
    for i:=range members{
        members[i]=i
    }
    var startIdx int
    var delta=k-1
    for n>1{
        //startIdx不用加1，因为会删除startIdx处元素，删除后下一个即变成startIdx处元素
        startIdx=(startIdx+delta)%len(members)
        //删除idx处的元素
        members=append(members[:startIdx],members[startIdx+1:]...)
        n--
    }
    return members[0]+1
}
