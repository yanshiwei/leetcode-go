class Solution {
    //给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
    //https://leetcode.cn/problems/permutations/solutions/9914/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/?orderBy=most_votes
public:
    void helper(vector<int>& nums,int start,vector<bool>&used,vector<int>&tmp,vector<vector<int>>&res){
        //结束条件
        if(start==nums.size()){
            res.push_back(tmp);
            return;
        }
        for(int i=0;i<nums.size();i++){
            if(used[i]){
                continue;
            }
            used[i]=true;
            tmp.push_back(nums[i]);
            helper(nums,start+1,used,tmp,res);
            tmp.pop_back();
            used[i]=false;
        }
    }
    vector<vector<int>> permute(vector<int>& nums) {
        if(nums.size()<1){
            return {};
        }
        vector<vector<int>> res;
        vector<int>tmp;
        vector<bool>used(nums.size(),false);//nums 中的所有整数 互不相同
        helper(nums,0,used,tmp,res);
        return res;
    }
};
