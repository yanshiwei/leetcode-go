class Solution {
    /*
    题目：
    给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
    题解：
    1 确定dp，dp[i][j],长度为i也就是以下标i-1为结尾的字符串s，和长度为 jj也就是以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]
    2 确定转移公示
    if s[i-1] == t[j-1]
        dp[i][j]=dp[i-1][j-1]+1
    else
        dp[i][j]=dp[i][j-1]
    3 初始化
        dp[0][j]=0，空串s与t下标在[0,j-1]中，不可能有相同⼦序列，所以全为0
        dp[i][0]=0,表示空串t与s下标在[0,i-1]中，不可能有相同⼦序列，所以全为0
    4 顺序，从左到右
    5 打印dp
    */
public:
    bool isSubsequence(string s, string t) {
        int slen=s.size();
        int tlen=t.size();
        // 如果要想s是t的子序列首先是t的长度要大于s的长度
        if(slen>tlen){
            return false;
        }
        vector<vector<int>> dp(slen + 1, vector<int>(tlen+ 1, 0));
        for(int i=1;i<=slen;i++){
            for(int j=1;j<=tlen;j++){
                if(s[i-1]==t[j-1]){
                    dp[i][j]=dp[i-1][j-1]+1;
                }else{
                    // 没有找到这需要t往前查找
                    dp[i][j]=dp[i][j-1];
                }
            }
        }
        // dp[i][j] 的最大值是否和s的大小一样长。
        if(dp[slen][tlen]==s.size()){
            return true;
        }
        return false;
    }
};
