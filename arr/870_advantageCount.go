func advantageCount(A []int, B []int) []int {
    /*
    田忌赛马。如果 A 中最小的牌 a 能击败 B 中最小的牌 b，那么我们应当将它们配对，因为此时在 A 中的每张牌都比 b 要大，所以不管我们在 b 前面放置哪张牌都可以得分。我们可以用手中最弱的牌来与 b 配对，这样会使 A 中剩余的牌严格地变大，因此会有更多得分点。否则， a 将无益于我们的比分，因为它无法击败任何牌。
    */
    var sortA=append([]int{},A...)
    var sortB=append([]int{},B...)
    sort.Ints(sortA)
    sort.Ints(sortB)
    var used=make(map[int][]int)//key is b value, and value is a value.存储的是大于b的所有A
    var remainList=make([]int,0)//丢弃着，顺序任意
    var j int
    for i:=range sortA{
        if sortA[i]>sortB[j]{
            used[sortB[j]]=append(used[sortB[j]],sortA[i])
            j++
        }else {
            remainList=append(remainList,sortA[i])
        }
    }
    var res []int
    for i:=range B {
        b:=B[i]
        beyondList:=used[b]
        var head int
        if len(beyondList)>0{
            head=beyondList[len(beyondList)-1]
            used[b]=beyondList[0:len(beyondList)-1]//pop
        }else{
            head=remainList[len(remainList)-1]
            remainList=remainList[0:len(remainList)-1]//pop
        }
        res=append(res,head)
    }
    return res
}
