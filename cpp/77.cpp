class Solution {
    /*
    给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
示例 1：
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
    */
public:
    void helper(int n,int k,int idx,vector<int>&tmp,vector<vector<int>>&res){
        if(tmp.size()==k){
            res.push_back(tmp);
            return;
        }   
        for(int i=idx;i<=n-(k-tmp.size())+1;i++){
            // old: i=n
            // new: i<=n-(k-tmp.size()) +1
            tmp.push_back(i);
            helper(n,k,i+1,tmp,res);
            tmp.pop_back();
        }
    }
    vector<vector<int>> combine(int n, int k) {
        if(n<1||k<1||n<k){
            return {};
        }
        vector<vector<int>> res;
        vector<int>tmp;
        helper(n,k,1,tmp,res);
        return res;
    }
};
