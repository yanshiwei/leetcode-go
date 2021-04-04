func kWeakestRows(mat [][]int, k int) []int {
    var res []int
    var m=len(mat)
    var n=len(mat[0])
    var sum []int
    for i:=0;i<m;i++{
        var oneRowSum int
        for j:=0;j<n;j++{
            oneRowSum+=mat[i][j]
        }
        res=append(res,i)
        sum=append(sum,oneRowSum)
    }
    //fmt.Println(sum)
    sort.Slice(res,func(i,j int)bool{
        if sum[res[i]]!=sum[res[j]]{
            return sum[res[i]]<sum[res[j]]
        }
        return res[i]<res[j]
    })
    return res[0:k]
}
