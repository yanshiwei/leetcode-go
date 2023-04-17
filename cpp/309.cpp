class Solution {
    /*
    题目：
    给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
    题解：
    //buy---buy---sel---frize---sel---sel---buy...
    dp[i][j]表示i天状态为j时所拥有的最大现金,j取个值：
    j=0表示买入；j=1卖出；j=3冷冻
    转移公式：
    dp[i][2]=dp[i-1][1]
    dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i],dp[i][2])//要么保持前一天状态（其中前一天状态要么是买出，要么是卖出后的冷冻），要么前一天还是买入当天卖出
    dp[i][0]=max(dp[i-1][0],dp[i-1][2]-prices[i])//要么保持前一天状态（其中前一天只能是买入），要么前一天还是冷冻当天卖出
    初始化：
    dp[0][0]=-prices[0]
    dp[0][1]=0;
    dp[0][2]=0;
    降维度：
    dp[2]=dp[1];
    dp[1]=max(dp[1],dp[0]+prices[i],dp[2])
    dp[0]=max(dp[0],dp[2]-prices[i])
    */
public:
    int maxProfit(vector<int>& prices) {
        vector<int>dp(3,0);
        dp[0]=-prices[0];
        for(int i=1;i<prices.size();i++){
            int temp=dp[1];
            dp[0]=max(dp[0],dp[2]-prices[i]);
            dp[1]=max(max(dp[1],dp[0]+prices[i]),dp[2]);
            dp[2]=temp;//冷冻期要使用上一阶段卖出后结果不能使用更新后的
        }
        return max(dp[1],dp[2]);
    }
};
