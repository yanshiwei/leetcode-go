class Solution {
    /*
    给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
    */
public:
    vector<vector<int>> generateMatrix(int n) {
        if(n<1){
            return {};
        }
        vector<vector<int>>res(n,vector<int>(n));
        int top=0;
        int bottom=n-1;
        int left=0;
        int right=n-1;
        int cnt=1;
        while(1){
            // top
            for(int j=left;j<=right;j++){
                res[top][j]=cnt++;
            }
            top++;
            if(top>bottom){
                break;
            }
            // right
            for(int i=top;i<=bottom;i++){
                res[i][right]=cnt++;
            }
            right--;
            if(left>right){
                break;
            }
            // bottom
            for(int j=right;j>=left;j--){
                res[bottom][j]=cnt++;
            }
            bottom--;
            if(top>bottom){
                break;
            }
            // left
            for(int i=bottom;i>=top;i--){
                res[i][left]=cnt++;
            }
            left++;
            if(left>right){
                break;
            }
        }
        return res;
    }
};
