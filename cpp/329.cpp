class Solution {
    /*
    给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
    */
public:
    int dfs(int x,int y,vector<vector<int>>& matrix,vector<vector<int>>&visited){
        if(visited[x][y]>0){
            return visited[x][y];
        }
        int maxV=1;
        int m=matrix.size();
        int n=matrix[0].size();
        for(int ii=0;ii<4;ii++){
            int tx=x+dx[ii];
            int ty=y+dy[ii];
            if(tx>=0&&tx<m&&ty>=0&&ty<n){
                if(matrix[tx][ty]>matrix[x][y])
                    maxV=max(maxV,1+dfs(tx,ty,matrix,visited));
            }
        }
        visited[x][y]=max(visited[x][y],maxV);
        return visited[x][y];
    }
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        int m=matrix.size();
        if(m<1){
            return 0;
        }
        int n=matrix[0].size();
        if(n<1){
            return 0;
        }
        //visited有两个作用：1.判断是否访问过，2.存储当前格子的最长递增长度
        vector<vector<int>>visited(m,vector<int>(n,0));
        int res=0;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                res=max(res,dfs(i,j,matrix,visited));
            }
        }
        return res;
    }
private:
    const vector<int> dx = {1, -1, 0, 0};
    const vector<int> dy = {0, 0, 1, -1};
};
