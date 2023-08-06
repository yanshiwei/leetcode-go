class Solution {
    /*
    n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
    */
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
    void helper(vector<string>&queens,int curRow,int &res){
        if(curRow==queens.size()){
            res++;
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
    int totalNQueens(int n) {
        if(n==1){
            return 1;
        }
        int res=0;
        vector<string>queens(n,string(n,'.'));
        helper(queens,0,res);
        return res;
    }
};
