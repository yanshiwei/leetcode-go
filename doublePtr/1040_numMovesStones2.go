func numMovesStonesII(stones []int) []int {
    /*
    移动首尾端点石子的位置，把它插入到数轴上未占用的点上，来使该石子不再是端点石子。
    参考：https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/solution/chua-dong-chuang-kou-by-xiaoneng-2/

    如1,100,102,105,200
    对于最大移动次数：
        我们将首端点石子1移动到[100,200]中的非占用点，这样就可得到一个结果，stones[n-1]-stones[1]-(n-2),之所以减去（n-2），是因为[100,200]间包含了已经占有位置的n-2,去掉1和200
        我们将尾端点石子200移动到[1,105]中的非占用点，这样就可得到一个结果，stones[n-2]-stones[0]-(n-2)
        最终的最大移动次数为：max(stones[n-1]-stones[1]-n+2, stones[n-2]-stones[0]-n+2)
    对于最小移动次数：
        使用滑动窗口来记录最小移动次数，当窗口内的石子个数大于n时，我们需要缩小窗口。already_stone来统计        窗口内的石子数，剩下来的石子数n-already_stone全部用来移动。注意特殊情况就是类似3,4,5,6,10这种         前n-1个石子连续，第n个石子不连续，但是我们不能直接将10移到2，应该是3移到8，10移到7，需要两步完成。因为端点10不能移动到2否则他也是端点，不不符合题意。
        设想设有个尺子，上面有 n 个刻度点，我们用这个尺子在石子从最左边到最右边移动，每动一次都查看下在尺子范围内有 m个石子，那么要使这个区间填满，就需要移动 n-m次。有一种特例就是前n-1个石子连续，第n个石子不连续
    */
    sort.Ints(stones)
    var n=len(stones)
    var minV=n
    //双指针窗口
    var j int
    for i:=0;j<n;j++{
        //当前窗口小于长度n
        for (stones[j]-stones[i]+1>n){
            i++
        }
        //当前窗口石头数目
        alreadyStone:=j-i+1
        //特殊情况：前m-1个石子的顺序连续，最后一个石子不连续，需要移动2步
        if alreadyStone==n-1&&stones[j]-stones[i]+1==n-1{
            minV=min(minV,2)
        }else{
            minV=min(minV,n-alreadyStone)
        }
    }
    return []int{minV,max(stones[n-1]-stones[1]-n+2,stones[n-2]-stones[0]-n+2)}
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
