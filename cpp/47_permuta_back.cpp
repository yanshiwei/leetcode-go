class Solution {
public:
    void helper(vector<int>& nums,int start,vector<bool>&used,vector<int>&tmp,vector<vector<int>>&res){
        // 结束条件
        if(start==nums.size()){
            res.push_back(tmp);
            return;
        }
        for(int i=0;i<nums.size();i++){
            if(used[i]){
                continue;
            }
            if(i>0&&nums[i]==nums[i-1]&&used[i-1]==false){
                //若该一样数值之前用过，且重复，去重
                continue;
            }
            used[i]=true;
            tmp.push_back(nums[i]);
            helper(nums,start+1,used,tmp,res);
            tmp.pop_back();
            used[i]=false;
        }
    }
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>>res;
        vector<int>tmp;
        vector<bool>used(nums.size(),false);//nums对应下标，不能使用值，因为存在相同值
        sort(nums.begin(),nums.end());//方便去重
        helper(nums,0,used,tmp,res);
        return res;
    }
};
