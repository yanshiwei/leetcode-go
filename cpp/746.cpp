class Solution {
    /*
    题目：
    给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
请你计算并返回达到楼梯顶部的最低花费。
    */
public:
    int minCostClimbingStairs(vector<int>& cost) {
        //令dp[i]到达第i层需要的最少花费
        int n=cost.size();
        vector<int>dp(n+1);//dp[n]为顶层
        // 注意，顶峰为最后⼀个元素后⾯，但是这⾥的 dp[n-1] 还没到达顶峰
        dp[0]=0;
        dp[1]=0;//到达第0层和第1层不需要费用
        for(int i=2;i<=n;i++){
            dp[i]=min(dp[i-2]+cost[i-2],dp[i-1]+cost[i-1]);
        }
        return dp[n];
    }
};
