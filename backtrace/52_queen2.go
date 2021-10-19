var res int
func totalNQueens(n int) int {
    res=0
    var board=make([][]string,n)
    for i:=range board{
        board[i]=make([]string,n)
        for j:=range board[i]{
            board[i][j]="."
        }
    }
    dfs(n,0,board)
    return res
}

func dfs(n int,row int, board[][]string){
   if row==n{
       res++
       return
   }
   for col:=0;col<n;col++{
       if (isValid(row,col,board,n)){
           board[row][col]="Q"
           dfs(n,row+1,board)
           board[row][col]="."//back
       }
   }
}

func isValid(row,col int,board[][]string, n int)bool{
    // check col
    for i:=0;i<row;i++{
        if board[i][col]=="Q"{
            return false
        }
    }
    // check 45
    for i,j:=row-1,col-1;i>=0&&j>=0;i,j=i-1,j-1{
        if board[i][j]=="Q"{
            return false
        }
    }
    // check 135
    for i,j:=row-1,col+1;i>=0&&j<n;i,j=i-1,j+1{
        if board[i][j]=="Q"{
            return false
        }
    }
    // no need check row because is judged by row
    return true
}
