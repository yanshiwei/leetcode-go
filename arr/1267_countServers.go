func countServers(grid [][]int) int {
    var res int
    var m=len(grid)
    if m<1 {
        return res
    }
    var n=len(grid[0])
    if n<1{
        return res
    }
    var rowCnt=make([]int,m)//每行服务器个数
    var colCnt=make([]int,n)//每列服务器个数
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if grid[i][j]==1{
                rowCnt[i]++
                colCnt[j]++
            }
        }
    }
    //在计数完成之后，我们再次遍历所有的服务器，如果位于 (x, y) 的服务器满足 rowCnt[x] > 1（即第 x 行至少有一台其它服务器）或 colCnt[y] > 1（即第 y 列至少有一台其它服务器），那么它就能够与至少一台其它服务器进行通信
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if grid[i][j]==1&&(rowCnt[i]>1||colCnt[j]>1){
                res++
            }
        }
    }
    return res
}
