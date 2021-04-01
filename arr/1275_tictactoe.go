func tictactoe(moves [][]int) string {
    var matrix=make([][]int,3)
    for i:=range matrix{
        matrix[i]=make([]int,3)
    }
    for i:=range moves{
        oneMove:=moves[i]
        if i%2==0{
            matrix[oneMove[0]][oneMove[1]]=1
        }else{
            matrix[oneMove[0]][oneMove[1]]=-1
        }
    }
    //对角线判断
    sum1:=matrix[0][0]+matrix[1][1]+matrix[2][2]
    if sum1==3{
        return "A"
    }else if sum1==-3{
        return "B"
    }
    sum2:=matrix[0][2]+matrix[1][1]+matrix[2][0]
    if sum2==3{
        return "A"
    }else if sum2==-3{
        return "B"
    }
    //横 纵判断
    for i:=0;i<3;i++{
        sumR:=matrix[i][0]+matrix[i][1]+matrix[i][2]
        sumC:=matrix[0][i]+matrix[1][i]+matrix[2][i]
        if sumR==3 ||sumC==3{
            return "A"
        }else if sumR==-3||sumC==-3{
            return "B"
        }
    }
    if len(moves)<9 {
        return "Pending"
    }
    return "Draw"
}
