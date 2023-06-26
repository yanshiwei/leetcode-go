class Solution {
    /*
    题目：
    你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
    题解：
    类似198，只不过这里是循环数组，也就是第一个和元素不能和最后一个元素一起。
    dp[i]=max(dp[i-1],dp[i-2]+nums[i])
    那么就有2种方案：
    1. 选第一个，则后续只能走到n-1
    2. 选第二个，则后续可以走到n
    取两种情况的最大值
    */
public:
    int rob(vector<int>& nums) {
        if(nums.size()<1){
            return 0;
        }
        if(nums.size()==1){
            return nums[0];
        }
        if(nums.size()==2){
            return max(nums[0],nums[1]);
        }
        return max(rob198(nums,0,nums.size()-2),rob198(nums,1,nums.size()-1));
    }
private:
    int rob198(vector<int>&nums,int start,int end){
        if(start==end){
            return nums[start];
        }
        vector<int>dp(nums.size());
        dp[start]=nums[start];
        dp[start+1]=max(nums[start],nums[start+1]);
        for(int i=start+2;i<=end;i++){
            dp[i]=max(dp[i-1],dp[i-2]+nums[i]);
        }
        return dp[end];
    }
};
