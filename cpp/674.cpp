class Solution {
    // 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度
    // dp[i] 表示以 i 结尾（不⼀定要从0开始）的连续递增⼦序列的⻓度
    // dp[i]=dp[i-1]+1 if nums[i]>nums[i-1]
public:
    int findLengthOfLCIS(vector<int>& nums) {
        int n=nums.size();
        vector<int>dp(n,1);//以下标i为结尾的数组的连续递增的子序列长度最少也应该是1，即就是nums[i]这一个元素。所以dp[i]应该初始1
        int maxV=1;
        for(int i=1;i<n;i++){
            if(nums[i]>nums[i-1]){
                dp[i]=dp[i-1]+1;
            }
            maxV=max(maxV,dp[i]);
        }
        return maxV;
    }
};
