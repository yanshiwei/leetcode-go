class Solution {
    //https://leetcode.cn/problems/last-stone-weight-ii/solutions/805162/yi-pian-wen-zhang-chi-tou-bei-bao-wen-ti-5lfv/
public:
    int lastStoneWeightII(vector<int>& stones) {
        // 问题转化为：把一堆石头分成两堆,求两堆石头重量差最小值,要让差值小,两堆石头的重量都要接近sum/2
        // 将一堆stone放进最大容量为sum/2的背包,求放进去的石头的最大重量MaxWeight
        // 求出后最终答案即为sum-2*MaxWeight;
        // 故可以得到01背包问题，背包容量是sum/2,要放入物品（也就是石头数组元素）的重量是石头重量，价值也是石头重量
        //dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
        //vector<vector<bool>>dp(n+1,vector<bool>(sum+1,false));
        int sum=0;
        for(int s:stones){
            sum+=s;
        }
        int target=sum/2;
        vector<int>dp(target+1);//dp[j]容量j时候最大价值
        dp[0]=0;// 背包没有空间的时候，就相当于装0
        for(int i=1;i<=stones.size();i++){
            // https://blog.csdn.net/m0_46379909/article/details/119887413
            // 参考这里为什么内层逆序，保证推导第i层的j的时候第i层的j-w[i]还没被更新，此时dp[j-w[i]]存储的还是第i-1层的dp[j-w[i]]
            for(int j=target;j>=stones[i-1];j--){
                dp[j]=max(dp[j],dp[j-stones[i-1]]+stones[i-1]);//注意dp里第一个是从i=1开始，而nums第一个是从i=0开始，故dp[i][...]对应nums[i-1]
            }
        }
        return sum-2*dp[target];
    }
};
