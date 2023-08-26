class Solution {
    // 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
    //说明：每次只能向下或者向右移动一步。
    //dp[i][j]从出发点到此处的最小和
    //dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
    //直接用grid替代dp
public:
    int minPathSum(vector<vector<int>>& grid) {
        int m=grid.size();
        if(m<1){
            return 0;
        }
        int n=grid[0].size();
        if(n<1){
            return 0;
        }
        //dp[i][j]从出发点到此处的最小和
        //dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j];
        // 这里可以直接在grid上操作，所以不需要单独建dp
        // 第一行, 只从从左到右
        for(int j=1;j<n;j++){
            grid[0][j]+=grid[0][j-1];
        }
        // 第一列, 只从上到下
        for(int i=1;i<m;i++){
            grid[i][0]+=grid[i-1][0];
        }
        for(int i=1;i<m;i++){
            for(int j=1;j<n;j++){
                grid[i][j]+=min(grid[i-1][j],grid[i][j-1]);
            }
        }
        return grid[m-1][n-1];
    }
};
