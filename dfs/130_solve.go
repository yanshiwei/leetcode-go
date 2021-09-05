func solve(board [][]byte)  {
    var m=len(board)
    if m<1{
        return
    }
    var n=len(board[0])
    if n<1{
        return
    }
    /*
    区分与边界上连通的'O'与真正被'X'包围的'O'.
    1.将与边界上连通的'O'标记为'A'
    2.将真正被'X'包围的'O'标记为'X',将'A'还原为'O'.
    */
    //遍历所有字母，首先将所有边界上的'O'以及与其相通的'O'标记为'A'
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            var isBorder bool
            if i==0||i==m-1||j==0||j==n-1{
                isBorder=true
            }
            if isBorder&&board[i][j]=='O'{
                dfs(board,i,j)
            }
        }
    }
    //再次遍历所有字母，将'O'标记为'X'(即被包围的'O'),将所有的'A'还原为'O'
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            //surrounded by x
            if board[i][j]=='O'{
                board[i][j]='X'
            }
            if board[i][j]=='A'{
                //border,recovery
                board[i][j]='O'
            }
        }
    }
}
func dfs(board[][]byte,r,c int){
    if r<0||c<0||r>=len(board)||c>=len(board[0]){
        //not in board
        return
    }
    if board[r][c]!='O'{
        return
    }
    //将该字母标记为'A',并对其邻节点进行dfs
    board[r][c]='A'
    dfs(board,r-1,c)
    dfs(board,r+1,c)
    dfs(board,r,c-1)
    dfs(board,r,c+1)
}
