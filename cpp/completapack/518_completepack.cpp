class Solution {
    /*
    题目：
    给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。 
题目数据保证结果符合 32 位带符号整数。
    题解：
    https://leetcode.cn/problems/coin-change-ii/solutions/821592/gong-shui-san-xie-xiang-jie-wan-quan-bei-6hxv/
     硬币相当于我们的物品，每种硬币可以选择「无限次」，很自然的想到「完全背包」
     定义dp[i][j] 为考虑前 i件物品，凑成总和为 j的方案数量。i取值coins，重量是其值coins[i]，没有价值
     对于第i个数字，假如值为t，有如下选择：
     选0个：dp[i][j]=dp[i-1][j]；
     选1个：dp[i][j]=dp[i-1][j-t]
     选2个：dp[i][j]=dp[i-1][j-2*t];
     选k个：dp[i][j]=dp[i-1][j-k*t];
     故dp[i][j]=fun(dp[i-1][j],dp[i-1][j-k*t]),0<k*t<=j
     dp[i][j-t]=fun(dp[i-1][j-t],dp[i-1][j-(k+1)*t])
     所以：dp[i][j]=fun(dp[i-1][j],dp[i][j-t])
     初始化：
     dp[0][j]=0;//第0件物品也就是没有物品，方案数肯定0
     dp[i][0]=1;//凑总和0的方案数只有一种
     降维：
     dp[j]=func(dp[j],dp[j-t])，第i层的dp[j-t]
     // 本题的话求方案数，则v[i]=0
     dp[j]=(dp[j]+dp[j-t]),第i-1层的dp[j],第i层的dp[j-t]
     //组合数mhttps://blog.csdn.net/qq_43530796/article/details/125410353
     //不讲究顺序，先物品后总额
    */
public:
    int change(int amount, vector<int>& coins) {
        int n=coins.size();
        vector<int>dp(amount+1,0);//dp[i]代表凑足金额i时候组合数
        // init 
        dp[0]=1;//凑总和0的方案数只有一种
        // 完全背包问题不讲究顺序，先物品后背包
        for(int i=0;i<coins.size();i++){
            int coin=coins[i];
            // 完全背包，正序，正序保证了dp[i][j-t]已经更新的是第i层
            for(int j=coin;j<=amount;j++){
                dp[j]+=dp[j-coin];
            }
        }
        return dp[amount];
    }
};
