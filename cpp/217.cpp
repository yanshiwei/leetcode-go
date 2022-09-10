class Solution {
    /*
    给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
    */
public:
    bool containsDuplicate(vector<int>& nums) {
        int nLen=nums.size();
        if(nLen<2){
            return false;
        }
        unordered_set<int>s;
        for(int i=0;i<nums.size();i++){
            if(s.find(nums[i])!=s.end()){
                return true;
            }
            s.insert(nums[i]);
        }
        return false;
    }
};
