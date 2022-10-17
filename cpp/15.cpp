class Solution {
    /*
    给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
    */
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> res;
        if(nums.size()<3){
            return res;
        }
        // 1 sort
        sort(nums.begin(),nums.end());
        // 固定nums[i]从后面双指针相向寻找
        for(int i=0;i<nums.size();i++){
            if(nums[i]>0){
                // 由于排序 后面不可能满足相加为0 之间退出
                break;
            }
            if(i>0&&nums[i]==nums[i-1]){
                // 重复元素：跳过，避免出现重复解
                continue;
            }
            int left=i+1,right=nums.size()-1;
            while(left<right){
                if(nums[i]+nums[left]+nums[right]==0){
                    res.push_back({nums[i],nums[left],nums[right]});
                    //判断左界和右界是否和下一位置重复，去除重复解。并同时将L,R 移到下一位置，寻找新的解
                    while(left<right&&nums[left]==nums[left+1]){
                        left++;
                    }
                    while(left<right&&nums[right]==nums[right-1]){
                        right--;
                    }
                    left++;
                    right--;
                }
                if(left>=right){
                    break;
                }
                if(nums[i]+nums[left]+nums[right]>0){
                    // 和大于 0 范围和过大
                    right--;
                }else if(nums[i]+nums[left]+nums[right]<0){
                    // 和小于 0 范围和过小
                    left++;
                }
            }
        }
        return res;
    }
};
