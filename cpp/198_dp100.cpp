class Solution {
    /*
    你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
    dp[i]代表nums[0:i]为止抢到的最大金额，有2种情况
    1 在i-1刚抢到，为了避免告警，这家不抢，故保留
    2 在i-2抢到，i-1保留，故i开始抢
    dp[i]=max(dp[i-1],dp[i-2]+nums[i])
    初始化：
    dp[0]=nums[0]
    dp[1]=max(nums[0],nums[1])
    */
public:
    int rob(vector<int>& nums) {
        if(nums.size()<1){
            return 0;
        }
        if(nums.size()==1){
            return nums[0];
        }
        int n=nums.size();
        vector<int>dp(n);
        dp[0]=nums[0];
        dp[1]=max(nums[0],nums[1]);
        for(int i=2;i<n;i++){
            dp[i]=max(dp[i-1],dp[i-2]+nums[i]);
        }
        return dp[n-1];
    }
};
