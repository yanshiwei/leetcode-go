func isValidSudoku(board [][]byte) bool {
    var row=make([][]int,9)
    for i:=range row{
        row[i]=make([]int,9)
    }
    var col=make([][]int,9)
    for i:=range col{
        col[i]=make([]int,9)
    }
    var box=make([][]int,9)
    for i:=range box{
        box[i]=make([]int,9)
    }
    for i:=0;i<9;i++{
        for j:=0;j<9;j++{
            if board[i][j]!='.'{
                num:=getInt(board[i][j])-1
                boxIdx:=(i/3)*3+j/3
                if row[i][num]>0||col[j][num]>0||box[boxIdx][num]>0{
                    return false
                }else{
                    row[i][num]=1
                    col[j][num]=1
                    box[boxIdx][num]=1
                }
            }
        }
    }
    return true   
}
func getInt(a byte)int {
    switch a{
        case '1':
        return 1
        case '2':
        return 2
        case '3':
        return 3
        case '4':
        return 4
        case '5':
        return 5
        case '6':
        return 6
        case '7':
        return 7
        case '8':
        return 8
        case '9':
        return 9
        default:
        return 0
    }
}
