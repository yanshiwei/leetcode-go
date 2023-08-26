class Solution {
    const int int_max=0x7fffffff;
    /*
    题目：
    给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
    完全背包问题：dp[i][j]=fun(dp[i-1][j],dp[i-1][j-k*t]+k*v),0<k*t<=j
    转为dp[i][j]=fun(dp[i-1][j],dp[i][j-t]+v)
    降维dp[j]=func(dp[j],dp[j-t]+v)
    由于是求最值，跟顺序无关，所以直接先物品后总额
    总额是n，重量是每个数字完全平方，价值是数量v=1
    */
public:
    int numSquares(int n) {
        // 范围：[1,n],每个数字可以被使用无限次，求凑出目标值 n 所需要用到的是最少数字个数是多少。
        // 第i个数字，i属于[1,n],假如值为t,对应重量是t=i*i，价值是1
        // 完全背包问题。
        // dp[i][j]前 i个数字，凑出数字总和 j所需要用到的最少数字数量
        // 对于第i个数字，假如值为t，有如下选择：
        // 选0个数字i,dp[i][j]=dp[i-1][j]
        // 选1个数字i,dp[i][j]=dp[i-1][j-t]+1
        // 选k个数字i,dp[i][j]=dp[i-1][j-k*t]+k*1
        // 故dp[i][j]=min(dp[i-1][j],dp[i-1][j-k*t]+k*1),0<k<=j
        // 由此dp[i][j-t]=min(dp[i-1][j-t],dp[i-1][j-(k-1)*t]+(k-1)*1)
        // 故dp[i][j]=min(dp[i-1][j],dp[i][j-t]+1)
        // 转为一维度：dp[j]=min(dp[j],dp[j-t]+1)
        vector<int>dp(n+1,int_max);//dp[i]代表和为i时候完全平方数的最小数量
        // init
        dp[0]=0;//和为0的完全平方数量为0个
        // 完全背包问题不讲究顺序，先物品后背包
        for(int i=1;i*i<=n;i++){
            // 物品nums列表[1,2,...,x]，其中x*x<=n，i不从0，从1开始
            int w=i*i;//nums[i]的重量，价值v[i]=1
            //完全背包，正序，正序保证了dp[i][j-w]已经更新的是第i行结果
            for(int j=w;j<=n;j++){
                dp[j]=min(dp[j],dp[j-w]+1);
            }
        }
        return dp[n];
    }
};
