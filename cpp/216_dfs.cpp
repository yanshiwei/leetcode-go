class Solution {
    /*
    找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次 
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
    */
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>>res;
        vector<int>tmp;
        helper(k,n,0,1,tmp,res);
        return res;
    }
private:
    void helper(int k,int remain,int cnt,int start,vector<int>&tmp,vector<vector<int>>&res){
        if(cnt>k){
            return;
        }
        // cnt<=k
        if(remain<0){
            return;
        }
        if(cnt==k){
            if(remain==0){
                res.push_back(tmp);
            }
            return;
        }
        for(int i=start;i<10;i++){
            tmp.push_back(i);
            helper(k,remain-i,cnt+1,i+1,tmp,res);
            tmp.pop_back();
        }
    }
};
