class Solution {
    // 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
    // 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
    // 问总共有多少条不同的路径？
    // dp dp[i][j]从左上角到(i,j)路径数,dp[i][j]=dp[i-1][j]+dp[i][j-1]
public:
    int uniquePaths(int m, int n) {
        if(m<1){
            return 0;
        }
        if(n<1){
            return 0;
        }
        vector<vector<int>>dp(m,vector<int>(n));
        // 第一行，左上角只能从左到右，故左上角到此位置的路径数均为1
        for(int i=0;i<n;i++){
            dp[0][i]=1;
        }
        // 第一列，左上角只能从上到下，故左上角到此位置的路径数均为1
        for(int i=0;i<m;i++){
            dp[i][0]=1;
        }
        for(int i=1;i<m;i++){
            for(int j=1;j<n;j++){
                dp[i][j]=dp[i-1][j]+dp[i][j-1];
            }
        }
        return dp[m-1][n-1];
    }
};
