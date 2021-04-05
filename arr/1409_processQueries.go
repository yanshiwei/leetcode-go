func processQueries(queries []int, m int) []int {
    var P=make([]int,m)
    var res []int
    for i:=0;i<m;i++{
        P[i]=i+1
    }
    for i:=range queries{
        query:=queries[i]
        pos:=-1
        for j:=range P{
            if query==P[j]{
                pos=j
                break
            }
        }
        res=append(res,pos)
        copy(P[1:pos+1],P[0:pos])
        P[0]=query
    }
    return res
}
