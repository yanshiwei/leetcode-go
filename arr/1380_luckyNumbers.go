func luckyNumbers (matrix [][]int) []int {
    var res[]int
    var m=len(matrix)
    var n=len(matrix[0])
    var cal []Info
    for i:=0;i<m;i++{
        var curMin=Info{Data:matrix[i][0],X:i,Y:0}
        for j:=0;j<n;j++{
            if curMin.Data>matrix[i][j]{
                curMin=Info{Data:matrix[i][j],X:i,Y:j}
            }
        }
        cal=append(cal,curMin)
    }
    for i:=range cal{
        one:=cal[i]
        var curMax=matrix[0][one.Y]
        for i:=0;i<m;i++{
            if curMax<matrix[i][one.Y]{
                curMax=matrix[i][one.Y]
            }
        }
        if curMax==one.Data{
            res=append(res,curMax)
        }
    }
    return res
}

type Info struct {
    Data int
    X int
    Y int
}
