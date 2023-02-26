class Solution {
    /*
    题目：
    给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
    题解：
    回溯经典总结：https://leetcode.cn/problems/palindrome-partitioning/solutions/640028/hui-su-fa-si-lu-yu-mo-ban-by-fuxuemingzh-azhz/
    */
public:
    bool isPalindrome(string s){
        if(s.empty()){
            return true;
        }
        int i=0,j=s.size()-1;
        while(i<=j){
            if(s[i]!=s[j]){
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
    //未探索区域：剩余的未搜索的字符串 s
    void dfs(string s,vector<vector<string>>&res,vector<string>path){
        //未探索区域满足结束条件
        if(s.empty()){
            res.push_back(path);
            return;
        }
        //for 选择 in 未探索区域当前可能的选择,每次选择可以选取 s 的 1 - length 个字符
        for(int i=1;i<=s.size();i++){
            string pre=s.substr(0,i);
            //if 当前选择符合要求:
            if(isPalindrome(pre)){
                path.push_back(pre);
                dfs(s.substr(i),res,path);//新的未探索区域：s 去除掉 pre 的剩余字符串
                path.pop_back();
            }
        }
    }
    vector<vector<string>> partition(string s) {
        vector<vector<string>>res;
        dfs(s,res,{});
        return res;
    }
};
