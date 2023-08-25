class Solution {
    /*
    给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
    */
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> res(numRows,vector<int>(0,0));
        for(int i=0;i<numRows;i++){
            for(int j=0;j<=i;j++){
                if(j==0||j==i){
                    res[i].push_back(1);
                }else{
                    res[i].push_back(res[i-1][j-1]+res[i-1][j]);
                }
            }
        }
        return res;
    }
};
