class Solution {
    /*
    一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。
    */
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        int m=obstacleGrid.size();
        if(m<1){
            return 0;
        }
        int n=obstacleGrid[0].size();
        if(n<1){
            return 0;
        }
        if(obstacleGrid[0][0]==1){
            // 起点就障碍物
            return 0;
        }
        vector<vector<int>>dp(m,vector<int>(n));
        dp[0][0]=1;
        // 第一行，只能从左到右 只有一种路径到此地
        for(int i=1;i<n;i++){
            if(obstacleGrid[0][i]!=1&&dp[0][i-1]){
                // 无障碍 且前一个可以到达
                dp[0][i]=1;
            }
        }
        // 第一列，只能从上到下 只有一种路径到此地
        for(int i=1;i<m;i++){
            if(obstacleGrid[i][0]!=1&&dp[i-1][0]){
                dp[i][0]=1;
            }
        }
        for(int i=1;i<m;i++){
            for(int j=1;j<n;j++){
                if(obstacleGrid[i][j]!=1){
                    dp[i][j]=dp[i-1][j]+dp[i][j-1];
                }
            }
        }
        return dp[m-1][n-1];
    }
};
