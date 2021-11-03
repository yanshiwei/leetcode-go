func maxProfit(k int, prices []int) int {
    if len(prices)<1{
        return 0
    }
    if k>=len(prices){
        //蜕变为可以交易多次
        return maxProfit2(prices)
    }
    var n=len(prices)
    var global=make([][]int,n)//到达第i天可以最多进行j次交易,最好的利润是多少
    var local=make([][]int,n)//到达第i天，最多可进行j次交易，并且最后一次交易在当天卖出的最好的利润是多少
    for i:=range global{
        global[i]=make([]int,k+1)
    }
    for i:=range local{
        local[i]=make([]int,k+1)
    }
    for i:=1;i<n;i++{
        var diff=prices[i]-prices[i-1]//第i天卖出
        for j:=1;j<k+1;j++{
            local[i][j]=max(global[i-1][j-1]+max(diff,0),local[i-1][j]+diff)//全局到i-1天进行j-1次交易(买入)，然后加上今天的交易（买入或卖出）；与local第i-1天j次交易（买入或卖出，最后一次交易是买入），然后加上今天的差值（今天卖出），按照要求第i天必须卖出所以第i-1天时已经达到了j次交易
            global[i][j]=max(global[i-1][j],local[i][j])//之所以是global[i-1][j]而非global[i-1][j-1]因为最后一次交易如果包含当天一定在局部最好的里面，否则一定在过往全局最优的里面
        }
    }
    return global[n-1][k]
}

func maxProfit2(prices []int)int{
    var res int
    for i:=1;i<len(prices);i++{
        if prices[i]>prices[i-1]{
            res+=prices[i]-prices[i-1]
        }
    }
    return res
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
