class Solution {
    /*
    给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
    */
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int res=0x7fffffff;
        int left=0,right=0;
        int sum=0;
        while(right<nums.size()){
            sum+=nums[right];
            while(sum>=target){
                // 最小长度
                res=min(res,right-left+1);
                sum-=nums[left];
                left++;
            }
            right++;
        }
        if(res==0x7fffffff){
            return 0;
        }
        return res;
    }
};
