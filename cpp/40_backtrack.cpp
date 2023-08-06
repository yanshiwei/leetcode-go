class Solution {
    /*
    给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。 
    dfs
    */
public:
    void helper(vector<int>& candidates, int target, int start,vector<int>&tmp,vector<vector<int>> &res){
        if(target<0){
            return;
        }
        if(target==0){
            res.push_back(tmp);
            return;
        }
        for(int i=start;i<candidates.size();i++){
            if(i>start&&candidates[i]==candidates[i-1]){
                // 去重
                continue;
            }
            tmp.push_back(candidates[i]);
            helper(candidates,target-candidates[i],i+1,tmp,res);
            tmp.pop_back();
        }
    }
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        if(candidates.size()<1){
            return {};
        }
        vector<vector<int>> res;
        vector<int>tmp;
        sort(candidates.begin(),candidates.end());// 去重关键
        helper(candidates,target,0,tmp,res);
        return res;
    }
};
