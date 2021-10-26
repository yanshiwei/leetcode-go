func countAndSay(n int) string {
    var res="1"
    for i:=2;i<=n;i++{
        var cur=""
        var m=len(res)
        var j int
        for j<m{
            k:=j+1
            for k<m&&res[j]==res[k]{
                k++
            }
            cnt:=k-j
            cur+=strconv.Itoa(cnt)+string(res[j])
            j=k
        }
        res=cur
    }
    return res
}
