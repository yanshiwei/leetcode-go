class Solution {
    /*
    题目：
    给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
    题解：
    定义一个三维数组 dp[n][k][2] 这里的n表示天数，其中：
    dp[i][j][0]：表示第i天交易了j次时卖出后的累计最大利润
    dp[i][j][1]：表示第i天交易了j次时买入后的累计最大利润
    第一次买入，要么前一天买入后保持，要么从初识态第一次买入
    dp[i][0][1]=max(dp[i-1][0][1],dp[i-1][0][0]-prices[i]);
    第一次卖出（完成一次交易），要么前一天卖出后保持，要么从前一天第一次买入后，当天卖出
    dp[i][1][0]=max(dp[i-1][1][0],dp[i-1][0][1]+prices[i])

    第二次买入：要么前一天买入后保持，要么是前一天第一次卖出后，当天就买入
    dp[i][1][1]=max(dp[i-1][1][1],dp[i-1][1][0]-prices[i])
    第二次卖出：要么前一天卖出后保持，要么是前一天第二次买入后，当天就卖出
    dp[i][2][0]=max(dp[i-1][2][0],dp[i-1][1][1]+prices[i])

    由此类推：
    第j次买入（还没有完成第j单故是j-1）：要么是从前一天第j-1次买入转换而来，当天保持；要么从前一天第j-1次卖出后当天买入
    dp[i][j-1][1]=max(dp[i-1][j-1][1],dp[i-1][j-1][0]-price[i])
    第j次卖出（卖出完成第j单故是j）：要么是从前一天第j次卖出转换而来，当天保持；要么从前一天第j次买入后当天卖出
    dp[i][j][0]=max(dp[i-1][j][0],dp[i-1][j-1][1]+prices[i])
    初始：
    for(int i = 0; i <= k; ++i){
        dp[0][i][0] = 0;// 第一天，不管交易多少次，卖出都是0收入
        dp[0][i][1] = -prices[0];// 第一天，不管交易多少次，买入的收入是-prices[0]
    }
    然后，我们对三维数组进行压缩，去掉最高维度的天数n，用二维数组dp[k][2]来代替。
这里的k就是交易了多少次，2表示买卖状态。
    第一次买入：dp[0][1]；第一次卖出：dp[1][0]
    第二次买入：dp[1][1]；第二次卖出：dp[2][0]
    第j次买入：dp[k-1][1]；第j次卖出：dp[j][0]
    递推公式：
    dp[j-1][1]=max(dp[j-1][1],dp[j-1][0]-prices[i])
    dp[j][0]=max(dp[j][0],dp[j-1][1]+prices[i])
    */
public:
    int maxProfit(int k, vector<int>& prices) {
        int n=prices.size();
        if(k>=n){
            // 相当于可以交易无数次，退化为题目：123
            return maxProfit2(prices);
        }
        vector<vector<int>>dp(k+1,vector<int>(2,0));//dp[j][z]交易第j次z=0时候卖出后的收益，dp[j][z]交易第j次z=1时候买入后的收益
        // init 根据三维降级而得,第一天
        for(int j=0;j<=k;j++){
            dp[j][0]=0;//第一天无论交易多少次，卖出后的收益都是0
            dp[j][1]=-prices[0];//第一天无论交易多少次，买入后的收益都是-price[0]
        }
        for(int i=0;i<n;i++){
            for(int j=1;j<=k;j++){
                dp[j-1][1]=max(dp[j-1][1],dp[j-1][0]-prices[i]);
                dp[j][0]=max(dp[j][0],dp[j-1][1]+prices[i]);
            }
        }
        return dp[k][0];//第k次交易的卖出
    }
    int maxProfit2(vector<int>&prices){
        // 可以交易无数次
        int res=0;
        for(int i=1;i<prices.size();i++){
            if(prices[i]>prices[i-1]){
                res+=prices[i]-prices[i-1];
            }
        }
        return res;
    }
};
