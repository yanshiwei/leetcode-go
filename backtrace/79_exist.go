func exist(board [][]byte, word string) bool {
    /*回溯*/
    for i:=0;i<len(board);i++ {
        for j:=0;j<len(board[0]);j++ {
            if helper(board,word,0,i,j) {
                // 从二维表格的每一个格子出发
                return true
            }
        }
    }
    return false
}

func helper(board [][]byte, word string,wordIdx int,x,y int)bool {
    if board[x][y]!=word[wordIdx]{
        // 当前位的字母不相等，此路不通
        return false
    }
    if wordIdx==len(word)-1 {
        // 最后一个字母也相等, 返回成功
        return true
    }
    tmp:=board[x][y]
    board[x][y]='0'//避免重复使用该位置
    wordIdx++
    up:=x>0&&helper(board,word,wordIdx,x-1,y)//上走
    left:=y>0&&helper(board,word,wordIdx,x,y-1)//左
    down:=x<len(board)-1&&helper(board,word,wordIdx,x+1,y)//下走
    right:=y<len(board[0])-1&&helper(board,word,wordIdx,x,y+1)//右
    if left||right||up||down {
        //中一条能走通，就算成功
        return true
    }
    board[x][y] = tmp; // 如果都不通，则回溯上一状态
    return false
}
