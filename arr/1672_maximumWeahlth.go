func maximumWealth(accounts [][]int) int {
    var maxV int
    for i:=0;i<len(accounts);i++{
        curSum:=0
        for j:=0;j<len(accounts[i]);j++{
            curSum+=accounts[i][j]
        }
        if curSum>maxV{
            maxV=curSum
        }
    }
    return maxV
}
