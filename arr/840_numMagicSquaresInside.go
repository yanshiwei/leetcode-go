func numMagicSquaresInside(grid [][]int) int {
    var res int
    var m=len(grid)
    if m<3{
        return res
    }
    var n=len(grid[0])
    if n<3 {
        return res
    }
    for r:=0;r<m-2;r++{
        for c:=0;c<n-2;c++{
            //事实是中心点肯定是5
            if grid[r+1][c+1]!=5 {
                continue
            }
            if isMagic([]int{grid[r][c],grid[r][c+1],grid[r][c+2],grid[r+1][c],grid[r+1][c+1],grid[r+1][c+2],grid[r+2][c],grid[r+2][c+1],grid[r+2][c+2]}){
                res++
            }
        }
    }
    return res
}

func isMagic(nums []int)bool{
    //grid[r][c],grid[r][c+1],grid[r][c+2]
    //grid[r+1][c],grid[r+1][c+1],grid[r+1][c+2]
    //grid[r+2][c],grid[r+2][c+1],grid[r+2][c+2]
    //ort.Ints(nums)
    for i:=range nums{
        if nums[i]<1 ||nums[i]>9{
            return false
        }
        if i>0&&nums[i]==nums[i-1]{
            //防止都是5
            return false
        }
    }
    if nums[0]+nums[1]+nums[2]!=15||nums[3]+nums[4]+nums[5]!=15||nums[6]+nums[7]+nums[8]!=15{
        return false
    }
    if nums[0]+nums[3]+nums[6]!=15||nums[1]+nums[4]+nums[7]!=15||nums[2]+nums[5]+nums[8]!=15{
        return false
    }
    if nums[0]+nums[4]+nums[8]!=15||nums[2]+nums[4]+nums[6]!=15{
        return false
    }
    return true
}
