func oddCells(m int, n int, indices [][]int) int {
    var arr=make([][]int,m)
    for i:=range arr{
        arr[i]=make([]int,n)
    }
    for i:=range indices{
        one:=indices[i]
        for c:=0;c<n;c++{
            arr[one[0]][c]+=1
        }
        for r:=0;r<m;r++{
            arr[r][one[1]]+=1
        }
    }
    var res int
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if arr[i][j]%2==1{
                res++
            }
        }
    }
    return res
}
