class Solution {
    /*
    题目：
    给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
    题解：
    dp[i][j] 代表 word1 到长度 i 转换成 word2 到 长度j 需要最少步数，i=0代表word1为空
    当word1[i] == word2[j]，dp[i][j]=dp[i-1][j-1]
    当word1[i] != word2[j]，dp[i][j]=min(dp[i-1][j-1],dp[i][j-1],dp[i-1][j])+1
        其中：dp[i-1][j-1] 表示替换操作；dp[i-1][j] 表示删除操作；dp[i][j-1] 表示添加操作
        对“dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。”的补充理解：
以 word1 为 "horse"，word2 为 "ros"，且 dp[5][3] 为例，即要将 word1的前 5 个字符转换为 word2的前 3 个字符，也就是将 horse 转换为 ros，因此有：
(1) dp[i-1][j-1]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 2 个字符 ro，然后将第五个字符 word1[4]（因为下标基数以 0 开始） 由 e 替换为 s（即替换为 word2 的第三个字符，word2[2]）
(2) dp[i][j-1]，即先将 word1 的前 5 个字符 horse 转换为 word2 的前 2 个字符 ro，然后在末尾补充一个 s，即插入操作
(3) dp[i-1][j]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 3 个字符 ros，然后删除 word1 的第 5 个字符
    */
public:
    int minDistance(string word1, string word2) {
        int n1=word1.size();
        int n2=word2.size();
        // dp[i][j] 代表 word1 到长度 i 转换成 word2 到 长度j 需要最少步数，i=0代表word1为空
        vector<vector<int>>dp(n1+1,vector<int>(n2+1,0));
        // 第一行,是 word1 为空变成 word2 最少步数，就是插入操作
        for(int j=1;j<=n2;j++){
            dp[0][j]=dp[0][j-1]+1;
        }
        // 第一列，是 word2 为空，需要的最少步数，就是删除操作
        for(int i=1;i<=n1;i++){
            dp[i][0]=dp[i-1][0]+1;
        }
        // i=0 j=0都初始化结束后，所以从1开始
        for(int i=1;i<=n1;i++){
            for(int j=1;j<=n2;j++){
                if(word1[i-1]==word2[j-1]){
                    dp[i][j]=dp[i-1][j-1];
                }else{
                    // 有三种操作：
                    // 增，word1的前i个字符编辑后只能对应到word2的前j-1个，故word1新增个--->dp[i][j-1]
                    // 删，word1的前i-1个字符编辑后能对应到word2的前j个，故word1删除last-->dp[i-1][j]
                    // 改，word1的前i-1个字符编辑后能对应到word2的前j-1个，故word1的最后一个更新-->dp[i-1][j-1]
                    dp[i][j]=min(dp[i][j-1],min(dp[i-1][j],dp[i-1][j-1]))+1;
                }
            }
        }
        return dp[n1][n2];
    }
};
