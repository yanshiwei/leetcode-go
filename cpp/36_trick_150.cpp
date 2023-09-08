class Solution {
    // https://leetcode.cn/problems/valid-sudoku/solutions/58669/36-jiu-an-zhao-cong-zuo-wang-you-cong-shang-wang-x/?envType=study-plan-v2&envId=top-interview-150
   // j/3+(i/3)*3的由来
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        // 存储每一行的每个数是否出现过
        vector<vector<int>>row(9,vector<int>(10,0));// 整个board有9行，第二维的维数10是为了让下标和数独中的数字对应
        // 存储每一列的每个数是否出现过
        vector<vector<int>>col(9,vector<int>(10,0));
        // 存储每一个box的每个数是否出现过，整个board有9个box
        vector<vector<int>>box(9,vector<int>(10,0));
        for(int i=0;i<9;i++){
            for(int j=0;j<9;j++){
                if(board[i][j]=='.'){
                    continue;
                }
                int num=board[i][j]-'0';
                if(row[i][num]){
                    // 行已经出现过
                    return false;
                }
                if(col[j][num]){
                    // 列已经出现过
                    return false;
                }
                int boxId=j/3+(i/3)*3;
                if(box[boxId][num]){
                    // box已经出现过
                    return false;
                }
                row[i][num]=1;
                col[j][num]=1;
                box[boxId][num]=1;
            }
        }
        return true;
    }
};
