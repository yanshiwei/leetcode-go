class Solution {
    /*
    题目：
    给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
    题解：
    https://leetcode.cn/problems/longest-increasing-subsequence/solutions/24173/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-2/?envType=study-plan-v2&envId=top-100-liked
    dp[i]以i结尾的子序列严格递增长度,注意 nums[i] 必须被选取
    与674题不同，这里子序列可以非连续
设 j∈[0,i)，考虑每轮计算新 dp[i]时，遍历 [0,i)列表区间，做以下判断： 
    1.当 nums[i]>nums[j]时： nums[i]可以接在 nums[j]之后（此题要求严格递增），此情况下最长上升子序列长度为 dp[j]+1 ； 
    2.当 nums[i]<=nums[j] 时： nums[i]无法接在 nums[j]之后，此情况上升子序列不成立，跳过。 
    上述所有 情况下计算出的 dp[j]+1的最大值，为直到 i的最长上升子序列长度（即 dp[i]）。实现方式为遍历 j 时，每轮执行 dp[i]=max(dp[i],dp[j]+1)。 转移方程： dp[i]=max(dp[i],dp[j]+1)for j in [0,i)
    初始化：dp[i]=1至少是nums[i]本身
    从左到右
    例如nums = [10,9,2,5,3,7,21,18]
          dp = [1,1,1,2,2,3,4,4]
        */
public:
    int lengthOfLIS(vector<int>& nums) {
        if(nums.size()<2){
            return nums.size();
        }
        int res=0;
        vector<int>dp(nums.size(),1);//dp[i]以i结尾的子序列严格递增长度,注意 nums[i] 必须被选取,所以dp[i]=1至少是nums[i]本身
        for(int i=1;i<nums.size();i++){
            //j∈[0,i),遍历j，每轮执行 dp[i]=max(dp[i],dp[j]+1) 
            // 求出[0,i)序列的最长不连续子序列长也就是dp[i]
            for(int j=0;j<i;j++){
                if(nums[i]>nums[j]){
                    dp[i]=max(dp[i],dp[j]+1);
                }
            }
            res=max(res,dp[i]);
        }
        return res;
    }
};
