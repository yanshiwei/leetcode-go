class Solution {
public:
    /*
    如果是普通数组，最大子数组和参考53题，https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-interview-150
    dp[i]=max(dp[i-1]+nums[i],nums[i])
    res=max(res,dp[i])
    循环数组分2种情况讨论：
    0 若数组均负数，直接返回最大值，不用求和；
    1 数组没有循环，则是连续数组，直接按照普通数组做法；数组若循环，包含首尾但不连续，可以先求中间数据的子数组的最小和（中间的仍然是连续的），然后整个数组总和减去最小值即为最大值，然后求这2个子情况下最大值
    故无论怎么样都需要求出中间数据也就是没有循环部分的最大子数组和，最小子数组和
    最大子数组和：
    dp[i]=max(dp[i-1]+nums[i],nums[i])
    res=max(res,dp[i])
    最小子数组和：
    dp[i]=min(dp[i-1]+nums[i],nums[i])
    res=min(res,dp[i])    
    题解：https://leetcode.cn/problems/maximum-sum-circular-subarray/solutions/2351107/mei-you-si-lu-yi-zhang-tu-miao-dong-pyth-ilqh/?envType=study-plan-v2&envId=top-interview-150
    */
    const int int_max=0x7fffffff;
    const int int_min=0x80000000;
    int maxSubarraySumCircular(vector<int>& nums) {
        int n=nums.size();
        if(n==1){
            return nums[0];
        }
        vector<int>maxDp(n);
        maxDp[0]=nums[0];
        vector<int>minDp(n);
        minDp[0]=nums[0];

        int maxv=int_min; // 求普通数组的最大子数组和
        int minv=int_max; // 求普通数组的最小子数组和
        int arr_sum=nums[0]; // 数组和-最小数组和，就是循环数组的最大子数组和
        int max_in_arr=nums[0]; // 判断数组是不是全是负数
        // 中间数据也就是没有循环部分的最大子数组和，最小子数组和
        for(int i=1;i<n;i++){
            max_in_arr=max(max_in_arr,nums[i]);

            maxDp[i]=max(maxDp[i-1]+nums[i],nums[i]);
            maxv=max(maxv,maxDp[i]);

            minDp[i]=min(minDp[i-1]+nums[i],nums[i]);
            minv=min(minv,minDp[i]);

            arr_sum+=nums[i];
        }
        // 若数组均负数，直接返回最大值，不用求和
        if(max_in_arr<0){
            return max_in_arr;
        }
        // 数组没有循环和有循环的最大值
        return max(maxv,arr_sum-minv);
    }
};
