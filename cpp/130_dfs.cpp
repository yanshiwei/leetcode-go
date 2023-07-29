class Solution {
    /*
    题目：
    给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
    题解：
    矩阵中有三种元素：
    字母 X；
    被字母 X 包围的字母 O；
    没有被字母 X 包围的字母 O。
本题要求将所有被字母 X 包围的字母 O都变为字母 X ，但很难判断哪些 O 是被包围的，哪些 O 不是被包围的。
注意到题目解释中提到：任何边界上的 O 都不会被填充为 X。 我们可以想到，所有的不被包围的 O 都直接或间接与边界上的 O 相连。
我们可以利用这个性质判断 O 是否在边界上，具体地说：
    对于每一个边界上的 O，我们以它为起点，标记所有与它直接或间接相连的字母 O；
    最后我们遍历这个矩阵，对于每一个字母：
    如果该字母被标记过，则该字母为没有被字母 X 包围的字母 O，我们将其还原为字母 O；
    如果该字母没有被标记过，则该字母为被字母 X 包围的字母 O，我们将其修改为字母 X。
    */
public:
    void solve(vector<vector<char>>& board){
        int m=board.size();
        if(m<1){
            return;
        }
        int n=board[0].size();
        if(n<1){
            return;
        }
        // 每一个边界上的 O，我们以它为起点，标记所有与它直接或间接相连的字母 O
        for(int i=0;i<m;i++){
            dfs(board,i,0);//第一列
            dfs(board,i,n-1);//最后一列
        }
        for(int j=1;j<n-1;j++){
            dfs(board,0,j);//第一行
            dfs(board,m-1,j);//最后一行
        }
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(board[i][j]=='A'){
                    // 该字母被标记过，则该字母为没有被字母 X 包围的字母 O
                    // 我们将其还原为字母 O；
                    board[i][j]='O';
                }else if(board[i][j]=='O'){
                    // 该字母没有被标记过，则该字母为被字母 X 包围的字母 O
                    // 按照题意我们将其修改为字母 X
                    board[i][j]='X';
                }
            }
        }
    }
    void dfs(vector<vector<char>>& board,int i,int j){
        // 边界判断
        if(i<0||j<0||i>=board.size()||j>=board[0].size()){
            return;
        }
        // 该位置是X或者访问过A
        if(board[i][j]!='O'){
            return;
        }
        //标志访问过
        board[i][j]='A';
        dfs(board,i-1,j);
        dfs(board,i+1,j);
        dfs(board,i,j-1);
        dfs(board,i,j+1);
    }
};
