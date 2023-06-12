class NumMatrix {
     /*定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：
NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
    // https://leetcode.cn/problems/range-sum-query-2d-immutable/solutions/629187/ru-he-qiu-er-wei-de-qian-zhui-he-yi-ji-y-6c21/
    */
public:
    vector<vector<int>>presum;
    // presum[i][j]代表(0,0)->(i-1,j-1)的和
    NumMatrix(vector<vector<int>>& matrix) {
        int m=matrix.size();
        if(m<1){
            return;
        }
        int n=matrix[0].size();
        if(n<1){
            return;
        }
        presum.resize(m+1);
        for(int i=0;i<m+1;i++){
            presum[i].resize(n+1);
        }
        for(int i=1;i<=m;i++){
            for(int j=1;j<=n;j++){
                presum[i][j]=presum[i][j-1]+presum[i-1][j]-presum[i-1][j-1]+matrix[i-1][j-1];
            }
        }
    }
    
    int sumRegion(int row1, int col1, int row2, int col2) {
        return presum[row2+1][col2+1]-presum[row2+1][col1]-presum[row1][col2+1]+presum[row1][col1];
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */
