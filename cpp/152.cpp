class Solution {
    // 题目：
    //给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//测试用例的答案是一个 32-位 整数。
//子数组 是数组的连续子序列。
    //dpMax[i]以i结尾的最大乘积，由于负负得正，故还需要存储dpMin[i]以i结尾的最小乘积
    //对于dpMax[i]：
    /* nums[i]>=0;
        若dpMax[i-1]>=0,则dpMax[i]=dpMax[i-1]*nums[i];
        若dpMax[i-1]<0,则dpMin[i]更小，则dpMax[i]=nums[i];
       nums[i]<0:
        若dpMax[i-1]>=0,此时dpMax[i-1]*nums[i]<dpMin[i-1]*nums[i]，则dpMax[i]=nums[i];
        若dpMax[i-1]<0,则dpMin[i-1]小0，则dpMax[i]=dpMin[i-1]*nums[i];
       总之：
       dpMax[i]=max(dpMin[i-1]*nums[i],max(dpMax[i-1]*nums[i],nums[i]))
    对于dpMin[i]:
     nums[i]>=0;
        若dpMin[i-1]>=0,则dpMin[i]=nums[i];
        若dpMin[i-1]<0，则dpMin[i]=dpMin[i-1]*nums[i];
     nums[i]<0:
        若dpMin[i-1]>=0,此时dpMax[i-1]>0，则dpMin[i]=dpMax[i-1]*nums[i];
        若dpMin[i-1]<0,则dpMin[i]=nums[i];
       总之：
       dpMin[i]=min(dpMin[i-1]*nums[i],min(dpMax[i-1]*nums[i],nums[i]))
    */
public:
    int maxProduct(vector<int>& nums) {
        int n=nums.size();
        vector<int>dpMax(n,0);
        dpMax[0]=nums[0];
        vector<int>dpMin(n,0);
        dpMin[0]=nums[0];  
        int res=nums[0];
        for(int i=1;i<n;i++){
            dpMax[i]=max(dpMin[i-1]*nums[i],max(dpMax[i-1]*nums[i],nums[i]));
            dpMin[i]=min(dpMin[i-1]*nums[i],min(dpMax[i-1]*nums[i],nums[i]));
            res=max(res,dpMax[i]);
        }      
        return res;
    }
};
