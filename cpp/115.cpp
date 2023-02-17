class Solution {
public:
/*
题目：
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
题解：
dp[i][j]代表s[i:]子序列中无差别存在t[j:]的数目
先考虑边界情况：
1. 当j=n时，t[j:]为空，空字符串是任何字符串的子序列，故dp[i][n]=1;
2. 当i=m时，s[i:]为空，非空字符串不是空字符串的子序列故dp[m][j]=0;
对于非空字符串，存在两种情况：
1. s[i]=t[j]:dp[i][j]由两种部分组成：
    1.1 如果s[i]与t[j]匹配，则考虑t[j+1:]作为s[i+1:]的子序列子序列数目dp[i+1][j+1];
    1.2 如果s[i]与t[j]不配，则考虑t[j:]作为s[i+1:]的子序列子序列数目dp[i+1][j];
因此当 s[i]=t[j]时，有 dp[i][j]=dp[i+1][j+1]+dp[i+1][j]；
2.s[i]!=t[j],s[i]与t[j]不配，则考虑t[j:]作为s[i+1:]的子序列子序列数目dp[i+1][j];
*/
    int numDistinct(string s, string t) {
        int m=s.size();
        int n=t.size();
        if(m<n){
            // 待匹配串必须不长与原始串
            return 0;
        }
        vector<vector<unsigned long long>>dp(m+1,vector<unsigned long long>(n+1,0));// 注意不能用int/long long否则溢出
        // init for 边界
        for(int i=1;i<=m;i++){
            dp[i][n]=1;
        }
        // 非边界处理
        for(int i=m-1;i>=0;i--){
            for(int j=n-1;j>=0;j--){
                if(s[i]==t[j]){
                    dp[i][j]=dp[i+1][j+1]+dp[i+1][j];
                }else{
                    dp[i][j]=dp[i+1][j];
                }
            }
        }
        return dp[0][0];
    }
};
