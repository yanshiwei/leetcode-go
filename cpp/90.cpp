class Solution {
    /*
    给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
    */
public:
    void helper(vector<int>&nums,int k, int start, vector<int>&tmp,vector<vector<int>>&res){
        if(tmp.size()==k){
            res.push_back(tmp);
            return;
        }
        for(int i=start;i<nums.size();i++){
            if(i>start&&nums[i]==nums[i-1]){
                continue;
            }
            tmp.push_back(nums[i]);
            helper(nums,k,i+1,tmp,res);
            tmp.pop_back();
        }
    }
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        vector<vector<int>> res;
        vector<int>tmp;
        sort(nums.begin(),nums.end());
        for(int i=0;i<=nums.size();i++){
            helper(nums,i,0,tmp,res);
        }
        return res;
    }
};
