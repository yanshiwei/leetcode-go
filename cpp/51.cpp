class Solution {
public:
    bool isValid(vector<string>&queens,int curRow,int col){
        // 当前行之前列不存在重复
        for(int i=0;i<curRow;i++){
            if(queens[i][col]=='Q'){
                return false;
            }
        }
        // 当前行 列之前对角线
        for(int i=curRow-1,j=col-1;i>=0&&j>=0;i--,j--){
            if(queens[i][j]=='Q'){
                return false;
            }
        }
        for(int i=curRow-1,j=col+1;i>=0&&j<queens.size();i--,j++){
            if(queens[i][j]=='Q'){
                return false;
            }
        }
        return true;       
    }
    void helper(vector<string>&queens,int curRow,vector<vector<string>> &res){
        if(curRow==queens.size()){
            res.push_back(queens);
            return;
        }
        // 当前行不同列尝试
        for(int i=0;i<queens.size();i++){
            if(isValid(queens,curRow,i)){
                queens[curRow][i]='Q';
                helper(queens,curRow+1,res);
                queens[curRow][i]='.';
            }
        }
    }
    vector<vector<string>> solveNQueens(int n) {
      vector<vector<string>> res;
      vector<string>queens(n,string(n,'.'));
      // 从0行开始递归
      helper(queens,0,res);
      return res;  
    }
};
