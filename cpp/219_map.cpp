class Solution {
    // 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        if(nums.size()<1||k<1){
            return false;
        }
        unordered_map<int,int>hash;
        for(int i=0;i<nums.size();i++){
            if(hash.count(nums[i])>0){
                if(i-hash[nums[i]]<=k){
                    return true;
                }
            }
            hash[nums[i]]=i;
        }
        return false;
    }
};
