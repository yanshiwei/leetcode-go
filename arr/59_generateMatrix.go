func generateMatrix(n int) [][]int {
    var left=0
    var right=n-1
    var top=0
    var bottom=n-1
    var totalNum=n*n
    var idx=1
    var matrix=make([][]int,n)
    for i:=range matrix{
        tmp:=make([]int,n)
        matrix[i]=tmp
    }
    for idx<=totalNum{
        //从左到右
        for i:=left;i<=right&&idx<=totalNum;i++{
            matrix[top][i]=idx
            idx++
        }
        top++
        //从上到下
        for i:=top;i<=bottom&&idx<=totalNum;i++{
            matrix[i][right]=idx
            idx++
        }
        right--
        //从右到左
        for i:=right;i>=left&&idx<=totalNum;i--{
            matrix[bottom][i]=idx
            idx++
        }
        bottom--
        //从下到上
        for i:=bottom;i>=top&&idx<=totalNum;i--{
            matrix[i][left]=idx
            idx++
        }
        left++
    }
    return matrix
}
