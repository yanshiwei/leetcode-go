class Solution {
    // https://leetcode.cn/problems/n-queens/solutions/399075/nhuang-hou-jing-dian-hui-su-suan-fa-tu-wen-xiang-j/
public:
    // 判断当前行 列 是否满足
    bool isValid(vector<string>&queens,int curRow,int col){
        // 相同列之前没有Q
        for(int i=0;i<curRow;i++){
            if(queens[i][col]=='Q'){
                return false;
            }
        }
        // 每一行只放一个Q，故在判断此行是否可以放时，次行之前还没放，故无需判断相同行
        // 当前行列 左对角线之前有没有Q
        for(int i=curRow-1,j=col-1;i>=0&&j>=0;i--,j--){
            if(queens[i][j]=='Q'){
                return false;
            }
        }
        // 当前行列 右对角线之前有没有Q
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
        for(int j=0;j<queens.size();j++){
            if(isValid(queens,curRow,j)){
                queens[curRow][j]='Q';
                helper(queens,curRow+1,res);
                queens[curRow][j]='.';
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
