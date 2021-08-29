var f bool
func solveSudoku(board [][]byte)  {
    /*
    DFS + 回溯
    需要注意的就是回溯的参数既要有横坐标也要有纵坐标
    还有就是重复数字的判断
    */
    f=false
    dfs(board,0,0,0)
}

func dfs(board[][]byte,row,col,finished int){
    var n=len(board)
    if finished==n*n{
        //都填完
        f=true
        return
    }
    if board[row][col]!='.'{
        if col<n-1{
            dfs(board,row,col+1,finished+1)
            return
        }
        //当col穷举到每一行的最后一个索引，即最后一列，就换到下一行(row+1)重新穷举
        dfs(board,row+1,0,finished+1)
        return
    }
    // 选择1到9
    for i:=1;i<=9;i++{
        if check(board,row,col,i){
            board[row][col]=getByte(i)
            if col<n-1{
                dfs(board,row,col+1,finished+1)
            }else{
                dfs(board,row+1,0,finished+1)
            }
            if f{
                return
            }
            // 撤销选择
            board[row][col]='.'
        }
    }
    return
}
func check(board [][]byte,row,col,num int)bool{
    var n=len(board)
    //判断同一列
    for i:=0;i<n;i++{
        if board[i][col]==getByte(num){
            return false
        }
    }
    //判断同一行
    for i:=0;i<9;i++{
        if board[row][i]==getByte(num){
            return false
        }
    }
    //判断同一块
    blockIdx:=(row/3)*3+col/3//记录当前小格子属于哪一大块，有0 ~ 8 共9大块
    blockStartX:=(blockIdx/3)*3// 记录大块的起始横坐标(对应board)
    blockStartY:=(blockIdx%3)*3// 记录大块的起始纵坐标
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            if board[blockStartX+i][blockStartY+j]==getByte(num){
                return false
            }
        }
    }
    return true
}

func getByte(num int)byte {
    switch num{
        case 1:
        return '1'
        case 2:
        return '2'
        case 3:
        return '3'
        case 4:
        return '4'
        case 5:
        return '5'
        case 6:
        return '6'
        case 7:
        return '7'
        case 8:
        return '8'
        case 9:
        return '9'
        default:
        return 0
    }
}
