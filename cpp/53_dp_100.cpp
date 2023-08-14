class Solution {
    /*
    给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
    */
public:
    int maxSubArray(vector<int>& nums) {
        if(nums.size()<1){
            return 0;
        }
        int res=0;
        int n=nums.size();
        vector<int>dp(n,0);//dp[i]以i结尾的连续子数组的最大和
        dp[0]=nums[0];
        res=dp[0];
        for(int i=1;i<n;i++){
            if(dp[i-1]<0){
                dp[i]=nums[i];
            }else{
                dp[i]=dp[i-1]+nums[i];
            }
            res=max(res,dp[i]);
        }
        return res;
    }
};
