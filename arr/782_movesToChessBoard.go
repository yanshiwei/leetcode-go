
func movesToChessboard(board [][]int) int {
    var m=len(board)
    if m<1 {
        return -1
    }
    var n=len(board[0])
    if n<1 {
        return -1
    }
    //行的目标风格有两种：10101010...或者010101...
    var row=make([]int,m)//行异同统计，每行与第一行是否相同相反，同则1 反则0
    //先从行来计算是否可能变换为棋盘。固定认为第一行是一种类型，其他行和第一行相同就是1，不相同但是相反就是0，如果不相同但是也不相反，肯定不能变换为棋盘，提前返回
    var firstRow=board[0]
    for i:=0;i<m;i++{
        curRow:=board[i]
        if isSame(firstRow,curRow){
            //相同
            row[i]=1
        }else if isReverse(firstRow,curRow){
            //相反
            row[i]=0
        }else{
            //不同但是不相反，肯定不能变换
            return -1
        }
    }
    //再计算行异同统计数组中1的数量是不是棋盘宽度的一半
    if isBalanceArr(row)==false{
        return -1
    }
    //接下来检查列，因为行已经满足了两种风格且互为逆序，因此列也肯定只有两种风格且互为逆序了。因为列也只要检查1的数量是否为棋盘宽度的一半就可以
    if isBalanceArr(board[0])==false{
        return -1
    }
    //计算行交换次数
    rowSwap:=getSwapFre(row)
    colSwap:=getSwapFre(board[0])
    return rowSwap+colSwap
}

func isSame(a,b []int)bool{
    for i:=range a{
        if a[i]!=b[i]{
            return false
        }
    }
    return true
}

func isReverse(a,b []int)bool {
    for i:=range a{
        if a[i]+b[i]!=1{
            return false
        }
    }
    return true
}

func getSwapFre(row []int)int{
    //行的目标风格有两种：10101010...或者010101...
    //计算出目前行的排列和目标风格有几个差异点，如果差异点为偶数个，则变化次数为差异点个数的一半（因为交换一次可以消除两个差异点），如果差异点为奇数个，则无法变换为目标风格（因为总有一个差异无法消除
    //取得两种风格的最小变换次数就是行变换棋盘的最小交换次数
    //101010...
    var fre int
    for i:=0;i<len(row);i++{
        if (i+1)%2==1{
            //奇数位置
            if row[i]!=1{
                fre++
            }
        }else{
            //偶数位置
            if row[i]!=0{
                fre++
            }
        }
    }
    if fre%2!=0{
        fre=-1
    }else {
        fre=fre/2
    }
    //010101...
    var anotherFre int
    for i:=0;i<len(row);i++{
        if (i+1)%2==1{
            //奇数位置
            if row[i]!=0{
                anotherFre++
            }
        }else{
            //偶数位置
            if row[i]!=1{
                anotherFre++
            }
        }
    }
    if anotherFre%2!=0{
        anotherFre=-1
    }else {
        anotherFre=anotherFre/2
    }
    if fre==-1 {
        return anotherFre
    }
    if anotherFre==-1{
        return fre
    }
    return min(fre,anotherFre)
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
//计算1的数目是否是棋盘宽度一半
func isBalanceArr(row []int)bool{
    var oneSum int
    for i:=range row {
        oneSum+=row[i]
    }
    if len(row)%2==0{
        //如果棋盘宽度为偶数，一半
        if oneSum!=len(row)/2{
            return false
        }
    }else{
        //如果棋盘宽度为奇数，则为/2（如010）或者/2+1（如101）
        if oneSum!=len(row)/2&&oneSum!=len(row)/2+1{
            return false
        }
    }
    return true
}
