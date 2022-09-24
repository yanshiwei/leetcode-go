class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
        /*
        给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
        */
        unordered_set<int>setx; // 可以去掉重复的行或者列
        unordered_set<int>sety;
        int m=matrix.size();
        int n=matrix[0].size();
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(matrix[i][j]==0){
                    setx.insert(i);
                    sety.insert(j);
                }
            }
        }
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(setx.count(i)){
                    matrix[i][j]=0;
                }
                if(sety.count(j)){
                    matrix[i][j]=0;
                }
            }
        }
    }
};
