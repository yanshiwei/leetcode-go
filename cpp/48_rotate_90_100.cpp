class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n=matrix.size();
        if(n<1){
            return;
        }
        //水平翻转,上半部分和下半部分交换  new[i][j]=old[n-row-1][col]
        for(int i=0;i<n/2;i++){
            for(int j=0;j<n;j++){
                swap(matrix[i][j],matrix[n-1-i][j]);
            }
        }
        //主对角线翻转，对角线左侧和右侧交换 new[i][j]=new[j][i]
        for(int i=0;i<n;i++){
            for(int j=0;j<i;j++){
                swap(matrix[i][j],matrix[j][i]);
            }
        }
    }
};
