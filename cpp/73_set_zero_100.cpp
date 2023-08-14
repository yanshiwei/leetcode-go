class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
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
        return;
    }
};
