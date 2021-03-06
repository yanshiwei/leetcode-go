func numRookCaptures(board [][]byte) int {
    //定义四个方向
    var dx=[]int{0,0,-1,1}
    var dy=[]int{-1,1,0,0}
    var m=len(board)
    var n=len(board[0])
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            //找到白车位置
            if board[i][j]=='R'{
                //分别判断四个方向
                var res int
                for k:=0;k<4;k++{
                    x:=i
                    y:=j
                    for {
                        x+=dx[k]
                        y+=dy[k]
                        if x<0||x>=m||y<0||y>=n||board[x][y]=='B'{
                            break
                        }
                        if board[x][y]=='p'{
                            res++
                            break//找到就停止
                        }
                    }
                }
                return res
            }
        }
    }
    return 0
}
