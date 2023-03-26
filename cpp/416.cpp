class Solution {
    /*
    题目：
    给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
    题解：
    给一个可装载重量为 sum / 2 的背包和 N 个物品，每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满
    01背包问题，https://blog.csdn.net/weixin_45746505/article/details/124543411
    */
public:
    // bool canPartition(vector<int>& nums) {
    //     int sum=0;
    //     for(int num:nums){
    //         sum+=num;
    //     }
    //     // 和为奇数时，不可能划分成两个和相等的集合
    //     if(sum%2==1){
    //         return false;
    //     }
    //     int n=nums.size();
    //     sum/=2;
    //     // 是否满足问题定义的是布尔DP
    //     // dp[i][j] = x 表示，对于前 i 个物品，当前背包的容量为 j 时，若 x 为 true，则说明可以恰好将背包装满，若 x 为 false，则说明不能恰好将背包装满
    //     vector<vector<bool>>dp(n+1,vector<bool>(sum+1,false));
    //     // init,dp[..][0] = true 和 dp[0][..] = false，因为背包没有空间的时候，就相当于装满了，而当没有物品可选择的时候，肯定没办法装满背包
    //     // 注意dp里第一个是从i=1开始，而nums第一个是从i=0开始，故dp[i][...]对应nums[i-1]
    //     for(int i=0;i<=n;i++){
    //         // 背包没有空间的时候，就相当于装满
    //         dp[i][0]=true;
    //     }
    //     for(int i=1;i<=n;i++){
    //         for(int j=1;j<=sum;j++){
    //             if(j<nums[i-1]){
    //                 // 背包容量不足，不能装入第 i 个物品
    //                 dp[i][j]=dp[i-1][j];
    //             }else{
    //                 // 装入或不装入背包
    //                 dp[i][j]=dp[i-1][j]||dp[i-1][j-nums[i-1]];
    //             }
    //         }
    //     }
    //     return dp[n][sum];
    // }
    bool canPartition(vector<int>& nums) {
        int sum=0;
        for(int num:nums){
            sum+=num;
        }
        if(sum%2==1){
            return false;
        }
        int n=nums.size();
        sum/=2;
        vector<bool> dp(sum + 1, false);
        dp[0]=true;
        for(int i=1;i<=n;i++){
            for(int j=sum;j>=nums[i-1];j--){
                dp[j] = dp[j] || dp[j - nums[i-1]];
            }
        }
        return dp[sum];
    }
};
