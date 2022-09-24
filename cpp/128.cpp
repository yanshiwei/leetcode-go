class Solution {
    /*
    给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
    */
public:
    int longestConsecutive(vector<int>& nums) {
        /*
        set:insert erase find count empty size swap lower_bound
        */
        int res=0;
        if(nums.size()<1){
            return res;
        }
        unordered_set<int> num_set;
        for(int i=0;i<nums.size();i++){
            num_set.insert(nums[i]);
        }
        // 若存在nums[i]-1说明nums[i]不是连续子序列的开始，跳过；若不存在，说明nums[i]是子序列开始，由此开始循环每次+1，直到不在set，维护此时最长长度
        for(int i=0;i<nums.size();i++){
            if(num_set.count(nums[i]-1)){
                // 存在nums[i]-1说明nums[i]不是连续子序列的开始，跳过
                continue;
            }
            // 若不存在，说明nums[i]是子序列开始，由此开始循环每次+1，直到不在set
            int curLen=0;
            int curV=nums[i];
            while(num_set.count(curV)){
                curLen++;
                curV++;
            }
            res=max(res,curLen);
        }
        return res;
    }
};
