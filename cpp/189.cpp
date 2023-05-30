class Solution {
    // 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
public:
    void rotate(vector<int>& nums, int k) {
        if(nums.size()<2){
            return;
        }
        vector<int>tmp(nums.size(),0);
        for(int i=0;i<nums.size();i++){
            tmp[(i+k)%nums.size()]=nums[i];
        }
        nums.swap(tmp);
        return;
    }
};
