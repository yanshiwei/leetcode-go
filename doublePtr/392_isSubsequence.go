func isSubsequence(s string, t string) bool {
    var n=len(s)
    var m=len(t)
    var i,j int
    for i<n&&j<m{
        if s[i]==t[j]{
            i++
        }
        j++
    }
    if i<n{
        return false
    }
    return true
}
