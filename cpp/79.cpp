class Solution {
    /*
    给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
    */
public:
    bool helper(vector<vector<char>>& board, string word,int i,int j,int idx){
        if(i<0||j<0||i>=board.size()||j>=board[0].size()){
            return false;
        }
        if(board[i][j]=='#'){
            //visted
            return false;
        }
        if(board[i][j]!=word[idx]){
            return false;
        }
        if(idx==word.size()-1){
            return true;
        }
        char t=board[i][j];
        board[i][j]='#';
        // four dir
        bool up=helper(board,word,i-1,j,idx+1);
        if(up){
            return true;
        }
        bool down=helper(board,word,i+1,j,idx+1);
        if(down){
            return true;
        }
        bool left=helper(board,word,i,j-1,idx+1);
        if(left){
            return true;
        }
        bool right=helper(board,word,i,j+1,idx+1);
        if(right){
            return true;
        }
        board[i][j]=t;
        return false;     
    }
    bool exist(vector<vector<char>>& board, string word) {
        if(board.size()<1){
            return false;
        }
        if(board[0].size()<1){
            return false;
        }
        for(int i=0;i<board.size();i++){
            for(int j=0;j<board[0].size();j++){
                if(helper(board,word,i,j,0)){
                    return true;
                }
            }
        }
        return false;
    }
};
