func shortestPathAllKeys(grid []string) int {
    var dx=[]int{-1,0,1,0}//左 下 右 上
    var dy=[]int{0,1,0,-1}
    var keyStatusList=[]int{1,2,4,8,16,32,64}//一共a-f共6个钥匙，在二进制分别对应
    //0000 0001--a(1)  0000 0010--b(2) 0000 0100--c(4) 0000 1000--d(8)
    //0001 0000--e(16)  0010 0000--f(32)
    //所有状态之和，1+2+4+8+16+32=63，故需要64
	visited := make([][][]bool, len(grid)) // i, j, keyState，状态和压缩最多1-63种状态,两外两位表示在当前钥匙状态下是否访问过该点。
    //之所以引入第三维，是因为不同钥匙状态 可以反复访问同一个位置
	for i := 0; i < len(grid); i++ {
		visited[i] = make([][]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = make([]bool, 64)
		}
	}
    var finishStatus=0//结束时状态
    var keyCnt int//钥匙数目
    var cacheList []Info
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++ {
            if grid[i][j]=='@' {
                //start
                cacheList=append(cacheList,Info{X:i,Y:j,CurKeyBit:0,StepCnt:0})
                visited[i][j][0]=true
            }
            if grid[i][j]>='a'&&grid[i][j]<='f' {
                //key
                keyCnt++
            }
        }
    }
    finishStatus=keyStatusList[keyCnt]-1//例如keyCnt=2则说明存在a b两个钥匙，则最后状态时0000 0011=3
    for len(cacheList)>0 {
        temp:=cacheList[0]
        cacheList=cacheList[1:]//pop
        if temp.CurKeyBit==finishStatus {
            //已经遍历完所有钥匙
            return temp.StepCnt
        }
        for i:=0;i<4;i++ {
            x:=dx[i]+temp.X
            y:=dy[i]+temp.Y
            if x<0||y<0||x>=len(grid)||y>=len(grid[0]){
                continue
            }
            if visited[x][y][temp.CurKeyBit] {
                //当前钥匙状态 已经访问过该节点
                continue
            }
            if grid[x][y]=='.'||grid[x][y]=='@' {
                //当遍历到空房间时，所拥有的钥匙状态不发生改变，故只需要更新一次标记数组的位置及路径数
                //起点也可以被通过
                visited[x][y][temp.CurKeyBit]=true
                cacheList=append(cacheList,Info{X:x,Y:y,CurKeyBit:temp.CurKeyBit,StepCnt:temp.StepCnt+1})
            }else if grid[x][y]>='a'&&grid[x][y]<='z' {
                visited[x][y][temp.CurKeyBit]=true
                //遇到钥匙，更新当前钥匙状态
                visited[x][y][temp.CurKeyBit|keyStatusList[grid[x][y]-'a']]=true//新的钥匙状态
                cacheList=append(cacheList,Info{X:x,Y:y,CurKeyBit:temp.CurKeyBit|keyStatusList[grid[x][y]-'a'],StepCnt:temp.StepCnt+1})
            }else if grid[x][y]>='A'&&grid[x][y]<='Z' &&temp.CurKeyBit&keyStatusList[grid[x][y]-'A']>0 {
                //遇到锁，首先判断有没有钥匙，与运算，求该bit位置与结果。
                //所拥有的钥匙状态不发生改变，故只需要更新一次标记数组的位置及路径数
                visited[x][y][temp.CurKeyBit]=true
                cacheList=append(cacheList,Info{X:x,Y:y,CurKeyBit:temp.CurKeyBit,StepCnt:temp.StepCnt+1})
            }
        }
    }
    return -1
}
type VisitedInfo struct{
    Visited [][]int
    Flag int
}
type Info struct {
    X int
    Y int
    CurKeyBit int
    StepCnt int
}
