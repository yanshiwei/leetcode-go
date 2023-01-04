class Solution {
    /*
    题目：
    给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
返回所有可能的结果。答案可以按 任意顺序 返回。
    https://leetcode.cn/problems/remove-invalid-parentheses/solutions/1070319/luo-ji-chai-jie-yi-ge-dfsde-xin-lu-li-ch-j97y/
    题目要求：删除几个 "(" 或 ")"，使得所有的左右括号都成对出现也就是通过删除m个左括号，n个右括号，使左右括号成对出现
    故第一步就是求出m/n.
    开始回溯：
    回溯推退出条件：
    1.判断是否已经不需要删除了，如果不需要删除了就进行合法性判定，合法就存储结果
    2.如果依旧需要删除，就逐个遍历字符:
        2.1 当前是'('且左括号的剩余删除次数还大于0，那就删掉这个左括号，并把左括号的剩余删除次数-1，然后递归调用
        2.2 当前是')'且右括号的剩余删除次数还大于0，那就删掉这个右括号，并把右括号的剩余删除次数-1，然后递归调用
    */
public:
    bool valid(string s){
        int leftBracketCnt=0;//左括号数
        for(auto &c:s){
            if(c=='('){
                leftBracketCnt++;
            }else if(c==')'){
                leftBracketCnt--;
                if(leftBracketCnt<0){
                    // 过程中一旦为负直接判为不满足条件
                    return false;
                }
            }
        }
        return leftBracketCnt==0;// 判断左右括号是否全部抵消
    }
    void helper(string s,int start,int m,int n,vector<string>&res){
        if(m==0&&n==0){
            //判断是否已经不需要删除了，如果不需要删除了就进行合法性判定，合法就存储结果
            if(valid(s)){
                res.push_back(s);
            }
            return;
        }
        // m>0||n>0,依旧需要删除，就逐个遍历字符:
        for(int i=start;i<s.size();i++){
            if(i>start&&s[i]==s[i-1]){
                //避免重复，形如"()())()"，我们删除第二个或第三个右括号，其实得到的结果都是一样的，也就是"()()()"
                continue;
            }
            if(s[i]=='('&&m>0){
                helper(s.substr(0,i)+s.substr(i+1),i,m-1,n,res);
            }
            if(s[i]==')'&&n>0){
                helper(s.substr(0,i)+s.substr(i+1),i,m,n-1,res);
            }
        }
    }
    vector<string> removeInvalidParentheses(string s) {
        // 求出m/n
        int m=0;// 代表要删除的左括号数量
        int n=0;// 代表要删除的右括号数量
        for(auto &c:s){
            if(c=='('){
                m++;
            }else if(c==')'&&m==0){
                // 遇右但无多余左，则代表要删一个右括号
                n++;
            }else if(c==')'&&m>0){
                // 遇右且有左，匹配消除左
                m--;
            }
        }
        // 最后剩下的m,n就是要删除的左右括号的最小数量
        vector<string> res;
        helper(s,0,m,n,res);
        return res;
    }
};
