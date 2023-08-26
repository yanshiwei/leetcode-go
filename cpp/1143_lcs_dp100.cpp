class Solution {
    /*
    题目：
    给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
    题解：
    https://leetcode.cn/problems/longest-common-subsequence/solutions/2287694/dong-tai-gui-hua-quan-wang-zui-qing-xi-l-02c8/?envType=study-plan-v2&envId=top-100-liked
    dp[i][j] 代表考虑 s1的前 i个字符、考虑 s2的前 j 的字符，形成的最长公共子序列长度
    s1[i]==s2[j]:dp[i][j]=dp[i-1][j-1]+1;
    s1[i]!=s2[j]:dp[i][j]=max(dp[i-1][j],dp[i][j-1]),代表必然不使用 s1[i]（但可能使用s2[j]）时 和 必然不使用 s2[j]（但可能使用s1[i]）时 LCS 的长度
    初始化：
    dp[0][j]或者dp[i][0]:s1空或者s2空，LCS肯定0
    */
public:
    int longestCommonSubsequence(string text1, string text2) {
        int m=text1.size();
        int n=text2.size();
        vector<vector<int>>dp(m+1,vector<int>(n+1,0));
        // init 
        // default is 0
        for(int i=1;i<=m;i++){
            for(int j=1;j<=n;j++){
                if(text1[i-1]==text2[j-1]){
                    dp[i][j]=dp[i-1][j-1]+1;
                }else{
                    // 代表必然不使用 s1[i]（但可能使用s2[j]）时;
                    // 必然不使用 s2[j]（但可能使用s1[i]）时 LCS 的长度
                    dp[i][j]=max(dp[i-1][j],dp[i][j-1]);
                }
            }
        }
        return dp[m][n];
    }
};
