func shiftGrid(grid [][]int, k int) [][]int {
    //k次操作后位置（i，j）移动
    //delteI=(j+k)/n,k有可能取值较大超过n，故还需取余(j+k)/n%m
    //deltaJ=(j+k)%n
    var m=len(grid)
    var n=len(grid[0])
    var ma=make([][]int,m)
    for i:=range ma{
        ma[i]=make([]int,n)
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            ma[(i+(j+k)/n)%m][(j+k)%n]=grid[i][j]
        }
    }
    return ma
}
