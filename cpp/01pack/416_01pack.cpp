class Solution {
    /*
    题目：
    给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
    题解：
    给一个可装载重量为 sum / 2 的背包和 N 个物品，每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满
    01背包问题，https://blog.csdn.net/weixin_45746505/article/details/124543411
    */
public:
// 优化前
    bool canPartition1(vector<int>& nums){
         // 要用背包问题，首先要确认
         // 1 背包容量，这里是sum/2;
         // 2 背包要放入的商品（也就是数组元素）的重量为元素的数值
         // 3 这里不是求最大价值和，而是判断能否放下，故func为或运算，价值忽略
         // dp[i][j] = x 表示，对于前 i 个物品，当前背包的容量为 j 时，若 x 为 true，则说明可以恰好将背包装满，若 x 为 false，则说明不能恰好将背包装满
         // vector<vector<bool>>dp(n+1,vector<bool>(sum+1,false));
         // 初始化，dp[..][0] = true 和 dp[0][..] = false，因为背包没有空间的时候，就相当于装满了，而当没有物品可选择的时候，肯定没办法装满背包
        if(nums.size()<1){
            return false;
        }
        int sum=accumulate(nums.begin(),nums.end(),0);
        if(sum%2==1){
            //不能平均
            return false;
        }
        int n=nums.size();
        sum/=2;
        vector<vector<bool>>dp(n+1,vector<bool>(sum+1,false));
        // 初始化
        for(int i=0;i<=n;i++){
            // 背包没有空间的时候，就相当于装满
            dp[i][0]=true;
        }
        for(int j=0;j<=sum;j++){
            // 没有物品可选择的时候，肯定没办法装满背包
            dp[0][j]=false;
        }
        // 注意dp里第一个是从i=1开始，而nums第一个是从i=0开始，故dp[i][...]对应nums[i-1]
        // 原始dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
        // 对于存在性问题，dp[i][j]=dp[i-1][j]||dp[i-1][j-w[i]]
        for(int i=1;i<=n;i++){
            // 01背包问题外循环从物品开始
            for(int j=1;j<=sum;j++){
                // 01背包问题内循环背包
                if(j<nums[i-1]){
                    dp[i][j]=dp[i-1][j];
                }else{
                    // 装入或不装入背包
                    dp[i][j]=dp[i-1][j]||dp[i-1][j-nums[i-1]];
                }
            }
        }
        return dp[n][sum];
    }
    bool canPartition(vector<int>& nums){
        return canPartition2(nums);
    }
// 优化后
    bool canPartition2(vector<int>& nums) {
        int sum=accumulate(nums.begin(),nums.end(),0);
        if(sum%2==1){
            return false;
        }
        int n=nums.size();
        sum/=2;
        vector<bool> dp(sum + 1, false);//dp[j]代表容量j时候是否能满
        dp[0]=true;// 背包没有空间的时候，就相当于装满
        // 注意dp里第一个是从i=1开始，而nums第一个是从i=0开始，故dp[i][...]对应nums[i-1]
        // 原始dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
        // 对于存在性问题，dp[i][j]=dp[i-1][j]||dp[i-1][j-w[i]]
        for(int i=1;i<=n;i++){
            // 降维度后，逆序，保证dp[j-w]是i-1层的
            for(int j=sum;j>=nums[i-1];j--){
                // https://blog.csdn.net/m0_46379909/article/details/119887413
                // 参考这里为什么内层逆序，保证推导第i层的j的时候第i层的j-w[i]还没被更新，此时dp[j-w[i]]存储的还是第i-1层的dp[j-w[i]]
                dp[j] = dp[j] || dp[j - nums[i-1]];
            }
        }
        return dp[sum];
    }
};
