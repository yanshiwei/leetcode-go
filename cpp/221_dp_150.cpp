class Solution {
    /*
    题目：
    在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
    题解：
    https://leetcode.cn/problems/maximal-square/solutions/253739/221-zui-da-zheng-fang-xing-tu-jie-shi-pin-yan-shi-/?envType=study-plan-v2&envId=top-interview-150
    当当前格=1，要形成非单1的正方形，以当前格为视角，则需要当前格、上、左、左上都是 1。
    然后取上、左、左上的边长最小的+1即为当前格子的边长；
    而上、左、左上的做法又可以类似上面，故可以使用动态规划
    以(i,j)(i,j)(i,j)为根据点最大的正方形的边长取决于红色蓝色绿色边长最短的那一个：
    其中dp[i−1][j−1]表示绿色的正方形，dp[i−1][j]表示红色的正方形，dp[i][j−1]表示蓝色的正方形
    */
public:
    int maximalSquare(vector<vector<char>>& matrix) {
        int n=matrix.size();
        if(n<1){
            return 0;
        }
        int m=matrix[0].size();
        if(m<1){
            return 0;
        }
        int res=0;//最大边长
        vector<vector<int>>dp(n,vector<int>(m,0));//dp[i][j]表示以第i行和j列为右下角的正方形的最大边长
        // init
        // 第一行，为0够不成正方形，为1就是当前格子
        for(int j=0;j<m;j++){
            if(matrix[0][j]=='1'){
                dp[0][j]=1;
                res=max(res,dp[0][j]);
            }
        }
        // 第一列，为0够不成正方形，为1就是当前格子
        for(int i=0;i<n;i++){
            if(matrix[i][0]=='1'){
                dp[i][0]=1;
                res=max(res,dp[i][0]);
            }
        }
        for(int i=1;i<n;i++){
            for(int j=1;j<m;j++){
                if(matrix[i][j]=='1'){
                    dp[i][j]=min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1]))+1;
                    res=max(res,dp[i][j]);
                }
            }
        }
        return res*res;
    }
};
