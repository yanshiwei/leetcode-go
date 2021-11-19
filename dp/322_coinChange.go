func coinChange(coins []int, amount int) int {
    if amount==0{
        return 0
    }
    var dp=make([]int,amount+1)//dp[i]表示当前所凑硬币总金额为i，至少所需的硬币数
    // init
    for i:=range dp{
        dp[i]=amount+1
    }
    dp[0]=0//i=0 dp=0
    min:=func(x,y int)int {
        if x<y{
            return x
        }
        return y
    }
    for i:=1;i<=amount;i++{
        for j:=0;j<len(coins);j++{
            if i-coins[j]>=0{
                dp[i]=min(dp[i],dp[i-coins[j]]+1)
                
            }else{
                // < 0 ignore
                continue
            }
        }
    }
    if dp[amount]==amount+1{
        return -1
    }
    return dp[amount]
}
const INTMAX=int(^uint(0)>>1)
