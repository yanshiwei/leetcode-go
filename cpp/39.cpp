class Solution {
    /*
    给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
对于给定的输入，保证和为 target 的不同组合数少于 150 个
示例 1：
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
    */
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<vector<int>> res;
        vector<int>tmp;
        helper(candidates,target,0,tmp,res);
        return res;
    }
    void helper(vector<int>& candidates, int target,int start,vector<int>&tmp,vector<vector<int>>&res){
        if(target<0){
            return;
        }
        if(target==0){
            res.push_back(tmp);
            return;
        }
        for(int i=start;i<candidates.size();i++){
            tmp.push_back(candidates[i]);
            helper(candidates,target-candidates[i],i,tmp,res);
            tmp.pop_back();
        }
    }
};
