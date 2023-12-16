class Solution {
    /*
    给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
    1.每种物品可以有无限多个 完全背包问题，dp[j]：凑足总额为j所需钱币的最少个数为dp[j], value是个数
    2. 完全背包问题：
    dp[i][j]=func(dp[i-1][j],dp[i][j-w[i]]+v[i])
    转为一维：dp[j]=func(dp[j],dp[j-w]+v),为力更新dp[i]是从小到大遍历 
    3. 组合不强调顺序、排列强调顺序！如果强调排列，我们应该先遍历背包，再遍历物品，只有这样才能让物品按一定顺序放入背包中。这里不强调顺序，就先物品后背包。也就是：
        外循环nums,内循环target,target正序且target>=nums[i];
        注意01背包上逆序，证明过程：https://mp.weixin.qq.com/s/nke1OjkhKACaONx1opk8AA
   4. 本题背包容量螫总金额，物品是不同额度硬币，硬币重量是额度，价值是个数，求硬币最少个数,func=min
        凑足总额为j - coins[i]的最少个数dp[j-coins[i]]，也就是
        加上一个钱币coins[i]也即dp[j-coins[i]]+1就是dp[j],其中0<j<=amount
        所以dp[j] 要取所有 dp[j - coins[i]] + 1 中最小的
        递推公式：dp[j] = min(dp[j - coins[i]] + 1, dp[j]);
    */
public:
    const int int_max=0x7FFFFFFF;
    int coinChange(vector<int>& coins, int amount) {
        vector<long>dp(amount+1,int_max);// long避免溢出
        // init,凑足总金额为0所需钱币的个数一定是0
        dp[0]=0;
        // 完全背包问题不讲究顺序，先物品后背包
        for(int i=0;i<coins.size();i++){
            // 先物品，物品重量是其额度，价值v=1
            int w=coins[i];
            for(int j=w;j<=amount;j++){
                // 后背包，背包容量是j
                // 从小到大排序，保证更新第i层
                dp[j]=min(dp[j],dp[j-w]+1);
            }
        }
        return dp[amount]==int_max?-1:dp[amount];// 有可能存在coins[2]，amount=3根本不能找到这样的组合
    }
};
