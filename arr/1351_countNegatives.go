func countNegatives(grid [][]int) int {
    var res int
    var m=len(grid)
    var n=len(grid[0])
    if grid[m-1][n-1]>=0{
        return res
    }
    for i:=m-1;i>=0;i--{
        for j:=n-1;j>=0;j--{
            if grid[i][j]>=0{
                break
            }else{
                res++
            }
        }
    }
    return res
}
