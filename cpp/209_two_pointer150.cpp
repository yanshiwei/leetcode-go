class Solution {
    /*
    给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
    */
public:
    const int int_max=0x7fffffff;
    int minSubArrayLen(int target, vector<int>& nums) {
        int res=int_max;
        int left=0,right=0;
        int sum=0;
        while(right<nums.size()){
            int c=nums[right];
            right++;
            // 窗口更新后操作
            sum+=c;
            // 窗口左移
            while(sum>=target){
                // sum>=target时，更新长度
                res=min(res,right-left);
                int d=nums[left];
                left++;
                // 窗口更新后操作
                sum-=d;
            }
        }
        return res==int_max?0:res;
    }
};
