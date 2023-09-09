class Solution {
    /*
    给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
    题解：
    https://leetcode.cn/problems/regular-expression-matching/solutions/296114/shou-hui-tu-jie-wo-tai-nan-liao-by-hyj8/
    */
public:
    bool isMatch(string s, string p) {
        //dp[i][j]长度j的p串是否匹配长度i的s
        int ls=s.size();
        int ps=p.size();
        if(ls<1||ps<1){
            return false;
        }
        vector<vector<bool>>dp(ls+1,vector<bool>(ps+1,false));
        // 特殊情况的初始化
        dp[0][0]=true;//空串自然匹配空串
        // s不为空串时，但p为空串处理肯定false
        // s为空串时，但p不为空串处理
        for(int j=1;j<=ps;j++){
            // 如果当前是*，则可以消掉一位上一位看上上一位
            if(p[j-1]=='*'){
                dp[0][j]=dp[0][j-2];
            }
        }
        for(int i=1;i<=ls;i++){
            for(int j=1;j<=ps;j++){
                if(s[i-1]==p[j-1]||p[j-1]=='.'){
                    //情况1：s[i-1]和 p[j-1] 是匹配的
                    dp[i][j]=dp[i-1][j-1];
                }else{
                    //s[i-1]和 p[j-1] 是不匹配的
                    if(p[j-1]!='*'){
                        //如果 p[j-1] 不是星号，那就真的不匹配
                        dp[i][j]=false;
                    }else{
                        // 如果 p[j-1]是星号,s[i-1]和 p[j-2] 是匹配的
                        if(s[i-1]==p[j-2]||p[j-2]=='.'){
                            //s[i-1]和 p[j-2] 是匹配的，具体有3种case
                            //case 1 p[j-1]=*让p[j-2]重复0次，这时结果就要看s[i-1]与p[j-3]
                            //case 2 p[j-1]=*让p[j-2]重复1次，这时结果就要看s[i-2]与p[j-3]
                            //case 2 p[j-1]=*让p[j-2]重复k次，拿出一个与s[i-1]抵消这时结果就要看s[i-2]与p[j-1]
                            dp[i][j]=dp[i][j-2]||dp[i-1][j-2]||dp[i-1][j];
                        }else {
                            // 如果 p[j-1]是星号,s[i-1]和 p[j-2] 是不匹配的
                            // p[j−1]星号可以干掉 p[j−2]，继续考察 s[i-1]和p[j-3]。
                            dp[i][j]=dp[i][j-2];
                        }
                    }
                }
            }
        }
        return dp[ls][ps];
    }
};
