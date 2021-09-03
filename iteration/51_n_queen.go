var res [][]string
var colFollow[][]string
func solveNQueens(n int) [][]string {
    res=nil
    colFollow=make([][]string,n)
    for i:=range colFollow{
        colFollow[i]=make([]string,n)
    }
    for i:=range colFollow{
        for j:=range colFollow[i]{
            colFollow[i][j]="."
        }
    }
    helper(n,0,colFollow)
    return res
}
func helper(n,row int,colFollow[][]string){
    if row==n{
        tmp:=make([]string,n)
        for i:=0;i<n;i++{
            tmp[i]=strings.Join(colFollow[i],"")
        }
        res=append(res,tmp)
        return
    }
    for col:=0;col<n;col++{
        if isValid(row,col,colFollow,n){
            colFollow[row][col]="Q"//放置
            helper(n,row+1,colFollow)
            colFollow[row][col]="."
        }
    }
}
func isValid(row,col int,colFollow[][]string,n int)bool{
    //check col
    for i:=0;i<col;i++{
        if colFollow[row][i]=="Q"{
            return false
        }
    }
    //check row
    for i:=0;i<row;i++{
        if colFollow[i][col]=="Q"{
            return false
        }
    }
    //check positive
    for i,j:=row-1,col-1;i>=0&&j>=0;i,j=i-1,j-1{
        if colFollow[i][j]=="Q"{
            return false
        }
    }
    //check negetive
    for i,j:=row-1,col+1;i>=0&&j<n;i,j=i-1,j+1{
        if colFollow[i][j]=="Q"{
            return false
        }
    }
    return true
}
