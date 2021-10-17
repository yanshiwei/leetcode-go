func titleToNumber(columnTitle string) int {
    var factor=1
    var res int
    for i:=len(columnTitle)-1;i>=0;i--{
        tmp:=columnTitle[i]-'A'+1
        res+=int(tmp)*factor
        factor*=26
    }
    return res
}
