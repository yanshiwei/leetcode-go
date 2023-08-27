class Solution {
    /*
    题目：
    给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
    题解：
    定义二维数组 dp[n][2]:
    1.dp[i][0] 表示第 i天卖出可获得的最大利润（卖出需要交易费）；
        dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i]-fee)
    2.p[i][1] 表示第 i 天买入可获得的最大利润(买入无需交易费)
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
    3 降维度
        dp[0]=max(dp[0],dp[1]+prices[i]-fee)
        dp[1]=max(dp[1],dp[0]-prices[i])
    */
public:
    int maxProfit(vector<int>& prices, int fee) {
        vector<int>dp(2,0);
        // init, first
        dp[0]=0;//第一天就卖出，利润0
        dp[1]=-prices[0];//第一天买入
        for(int i=1;i<prices.size();i++){
            dp[0]=max(dp[0],dp[1]+prices[i]-fee);
            dp[1]=max(dp[1],dp[0]-prices[i]);
        }
        return dp[0];
    }
};
