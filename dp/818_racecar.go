func racecar(target int) int {
    //无论正方向还是负方向，经过n个A可以移动（2^n）-1个距离
    //R操作是换方向
    //使用一个一维数组dp，dp[i]代表从0走到位置i处所需要的最小步数
    /*
    因为先向前走forward步再向后走back步与先向后走back步再向前走forward步最后到达的位置相同，所以可以假设永远保持第一步是向前走的。第一步有三种情况：
    1 刚好走forward步后到达了目标位置i，则dp[i] = forward;
    2 forward>i,这时需要再往回走，再加上回头的那一步，即AAAR，长度是forward+1，此时，我们相对i的距离是(2^forward)-1-i=j-i，dp[i] = Math.min(dp[i], forward + 1 + dp[j - i]);
    3 forward<i,这时候先反向行驶一会儿，然后再反向，驶向终点，这种模式是A……ARA……AR,长度是forward+1+back+1,此时我们相对i的距离是i-((2^forward-1)-(2^back-1))=i-j+k
    */
    var dp []int=make([]int,target+1)//0到其他位置的最小步数
    for i:=1;i<=target;i++ {
        dp[i]=INTMAX
        for forward:=1;(1<<forward)-1<2*i;forward++ {
            //第一步走forward上限是2i
            j:=(1<<forward)-1//forward pos
            if j==i {
                dp[i]=forward
            }else if j>i {
                dp[i]=min(dp[i],forward+1+dp[j-i])
            }else {
                for back:=0;back<forward;back++ {
                    //我们需要枚举可能的back，找到最小值
                    k:=(1<<back)-1//back pos
                    dp[i]=min(dp[i],forward+1+back+1+dp[i-j+k])
                }
            }
        }
    }
    return dp[target]
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
const INTMAX=int(^uint(0)>>1)
