class Solution {
public:
    // dp[i]是以s[i]为结尾的字串中，所形成的最长有效子串的长度
    int longestValidParentheses(string s) {
        if(s.size()<2){
            return 0;
        }
        int n=s.size();
        vector<int>dp(n,0);
        // init
        dp[0]=0;
        int res=0;
        for(int i=1;i<n;i++){
            if(s[i]==')'){
                // right bracket
                if(s[i-1]=='('){
                    //保底长度有2，且要看s[i-2]是什么情况
                    if(i-2>=0){
                        dp[i]=dp[i-2]+2;
                    }else{
                        dp[i]=2;
                    }
                }else{
                    //s[i-1]和s[i]是'))'，以s[i-1]为结尾的最长有效长度为dp[i-1]，跨过这个长度来看s[i-dp[i-1]-1]这个字符：
                    if(i-1-dp[i-1]>=0&&s[i-1-dp[i-1]]=='('){
                        //(dp[i-1])-->((...))
                        if(i-1-dp[i-1]-1>=0){
                            //...(dp[i-1])
                            dp[i]=dp[i-1-dp[i-1]-1]+dp[i-1]+2;
                        }else{
                            //(dp[i-1])
                            dp[i]=dp[i-1]+2;
                        }
                    }else{
                        //s[i-1-dp[i-1]]不存在或等于‘)'，则到s[i]是dp[i-1])不合法
                        dp[i]=0;
                    }
                }
            }else{
                // left bracket
                dp[i]=0;
            }
            res=max(res,dp[i]);
        }
        return res;
    }
};
