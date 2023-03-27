class Solution {
    /*
    给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
    */
public:
    const int int_max=0x7FFFFFFF;
    int coinChange(vector<int>& coins, int amount) {
        // 每种物品可以有无限多个 完全背包问题，dp[j]：凑足总额为j所需钱币的最少个数为dp[j], value是个数
        // 完全背包问题：外循环nums,内循环target,target正序且target>=nums[i];注意01背包上逆序，证明过程：https://mp.weixin.qq.com/s/nke1OjkhKACaONx1opk8AA
        // 凑足总额为j - coins[i]的最少个数dp[j-coins[i]]，也就是
        // 加上一个钱币coins[i]也即dp[j-coins[i]]+1就是dp[j],其中0<j<=amount
        // 所以dp[j] 要取所有 dp[j - coins[i]] + 1 中最小的
        // 递推公式：dp[j] = min(dp[j - coins[i]] + 1, dp[j]);
        vector<long>dp(amount+1,int_max);
        // init,凑足总金额为0所需钱币的个数一定是0
        dp[0]=0;
        for(int coin:coins){
            for(int j=coin;j<=amount;j++){
                dp[j]=min(dp[j],dp[j-coin]+1);
            }
        }
        return dp[amount]==int_max?-1:dp[amount];
    }
};
