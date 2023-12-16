class Solution {
public:
    /*
    令dp[i][j]为字符子串s1[0, i),s2[0, j)能否组成s3[0, i+j)。 
    ps: s1[0, i) = s1.substring(0, i),不包含i。
    对于dp[i][j]该状态来说，要想组成s3[0,i+j)，其s3[0, i+j)最后一个字符s3[i+j-1]要么来自s1[i-1], 要么来自s2[j-1],因此，状态转移：
    1. 若s1[i-1]==s3[i+j-1]:dp[i][j]=dp[i-1][j],i>0;
    2. 若s2[j-2]==s3[i+j-1]:dp[i][j]=dp[i][j-1],j>0;
    初始化dp[0][0]=true表示两个空字符串能够组成一个空字符串
    */

    bool isInterleave(string s1, string s2, string s3) {
        int n=s1.size();
        int m=s2.size();
        if(n+m!=s3.size()){
            return false;
        }
        vector<vector<bool>>dp(n+1,vector<bool>(m+1,false));
        for(int i=0;i<=n;i++){
            for(int j=0;j<=m;j++){
                if(i==0&&j==0){
                    dp[i][j]=true;//init
                }else{
                    if(i>0&&s1[i-1]==s3[i+j-1]){
                        dp[i][j]=(dp[i][j]||dp[i-1][j]);
                    }
                    if(j>0&&s2[j-1]==s3[i+j-1]){
                        dp[i][j]=(dp[i][j]||dp[i][j-1]);
                    }                    
                }
            }
        }
        return dp[n][m];
    }
};
