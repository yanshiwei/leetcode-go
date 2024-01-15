class Solution {
    /*
    给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
    */
public:
    //visited有两个作用：1.判断是否访问过，2.存储当前格子的最长递增长度
    int longestIncreasingPathCore(vector<vector<int>>& matrix,vector<vector<int>>&visited,int i,int j) {
        // 访问过判断
        if(visited[i][j]>0){
            return visited[i][j];
        }
        int maxV=1;//最大长度，默认最小是1，也就是本身
        int m=matrix.size();
        int n=matrix[0].size();
        for(int k=0;k<4;k++){
            int curi=i+dx[k];
            int curj=j+dy[k];
            if(curi>=0&&curi<m&&curj>=0&&curj<n) {
                if(matrix[curi][curj]>matrix[i][j]){
                    maxV=max(maxV,1+longestIncreasingPathCore(matrix,visited,curi,curj));
                }
            }
        }
        visited[i][j]=max(visited[i][j],maxV);
        return visited[i][j];
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
                res=max(res,longestIncreasingPathCore(matrix,visited,i,j));
            }
        }
        return res;
    }
private:
    const vector<int> dx = {1, -1, 0, 0};
    const vector<int> dy = {0, 0, 1, -1};
};
