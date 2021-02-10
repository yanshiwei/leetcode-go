func maxAreaOfIsland(grid [][]int) int {
    var res int
    var m=len(grid)
    if m<1 {
        return res
    }
    var n=len(grid[0])
    if n<1 {
        return res
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if grid[i][j]==1{
                curArea:=dfsArea(grid,i,j)
                res=max(curArea,res)
            }
        }
    }
    return res
}

func dfsArea(grid[][]int,row,col int)int{
    if row<0||row>=len(grid){
        return 0
    }
    if col<0||col>=len(grid[0]){
        return 0
    }
    if grid[row][col]!=1{
        return 0
    }
    grid[row][col]=2//标记已经访问
    left:=dfsArea(grid,row-1,col)
    right:=dfsArea(grid,row+1,col)
    up:=dfsArea(grid,row,col-1)
    down:=dfsArea(grid,row,col+1)
    return 1+left+right+up+down       
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
