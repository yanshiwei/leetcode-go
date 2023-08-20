class Solution {
    /*
    题目：
    题解：
    https://leetcode.cn/problems/palindrome-partitioning/solutions/54233/hui-su-you-hua-jia-liao-dong-tai-gui-hua-by-liweiw/?envType=study-plan-v2&envId=top-100-liked
    产生前缀字符串的时候，判断前缀字符串是否是回文：
    1. 如果前缀字符串是回文，则可以产生分支和结点
    2. 如果前缀字符串不是回文，则不产生分支和结点，这一步是剪枝操作
    在叶子结点是空字符串的时候结束，此时 从根结点到叶子结点的路径，就是结果集里的一个结果，使用深度优先遍历，记录下所有可能的结果
    */
public:
    bool isPalindrome(string s){
        if(s.empty()){
            return true;
        }
        int i=0;
        int j=s.size()-1;
        while(i<=j){
            if(s[i]!=s[j]){
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
    void helper(string s, vector<vector<string>>&res,vector<string>path){
        // 在叶子结点是空字符串的时候结束，此时 从根结点到叶子结点的路径，就是结果集里的一个结果，使用深度优先遍历，记录下所有可能的结果
        //未探索区域满足结束条件
        if(s.empty()){
            res.push_back(path);
            return;
        }
        //for 选择 in 未探索区域当前可能的选择,每次选择可以选取 s 的 1 - length 个字符
        for(int i=1;i<=s.size();i++){
            // 截取的前一个字符串
            string select_str=s.substr(0,i);
            if(isPalindrome(select_str)){
                path.push_back(select_str);
                helper(s.substr(i),res,path);//新的未探索区域：s 去除掉 pre 的剩余字符串
                path.pop_back();
            }
        }
    }
    vector<vector<string>> partition(string s) {
        vector<vector<string>> res;
        if(s.size()<1){
            return res;
        }
        vector<string>path;
        helper(s,res,path);
        return res;
    }
};
