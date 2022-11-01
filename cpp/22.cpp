class Solution {
    /*
    数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
    */
public:
    vector<string> generateParenthesis(int n) {
        if(n<1){
            return {};
        }
        vector<string> res;
        helper(n,n,"",res);
        return res;
    }
    void helper(int left,int right,string out,vector<string>&res){
        // 退出条件 剩余左括号右括号数目为0
        if(left==0&&right==0){
            res.push_back(out);
            return;
        }
        if(left>0){
            //只要(有剩，就可以选(。 (((((这么选，都还不能判定为非法
            helper(left-1,right,out+"(",res);
        }
        if(left<right){
            //当剩下的)比(多时即使用的(多于)时可以选择)，否则，)不能选，选了就非法，因为：剩下的)比(少，即，使用的)比(多，不能成双成对。
            helper(left,right-1,out+")",res);
        }
    }
};
