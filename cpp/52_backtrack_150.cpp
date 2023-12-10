class Solution {
    /*
    n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
相比较51题：https://leetcode.cn/problems/n-queens/description/?envType=study-plan-v2&envId=top-100-liked
本题求的是方案数量，51题求的是所有不同的 n 皇后问题 的解决方案
https://leetcode.cn/problems/n-queens/solutions/399075/nhuang-hou-jing-dian-hui-su-suan-fa-tu-wen-xiang-j/
    DFS
    */
public:
    bool isValid(vector<string>&queens, int cur_row, int col){
        // 当前行之前列不存在重复
        for(int i=0;i<cur_row;i++){
            if(queens[i][col]=='Q'){
                return false;
            }
        }
        // 当前行左对角线之前有没有Q
        for(int i=cur_row-1,j=col-1;i>=0&&j>=0;i--,j--){
            if(queens[i][j]=='Q'){
                return false;
            }
        }
        // 当前行右对角线之前有没有Q
        for(int i=cur_row-1,j=col+1;i>=0&&j<queens.size();i--,j++){
            if(queens[i][j]=='Q'){
                return false;
            }
        }
        return true;
    }
    void helper(vector<string>&queens, int cur_row, int&res){
        if(cur_row==queens.size()){
            // 可以走到最后一行
            res++;
            return;
        }
        // 当前行，不同列去尝试，由于是一行行递归，天然保证了同一行不会有冲突，故只需判断列和对角
        for(int j=0;j<queens.size();j++){
            if(isValid(queens,cur_row,j)){
                // 当前位置可以放
                queens[cur_row][j]='Q';
                // 递归下一行
                helper(queens,cur_row+1,res);
                queens[cur_row][j]='.';// 方便下一个位置继续处理
            }
        }
    }
    int totalNQueens(int n) {
        if(n==1){
            return 1;
        }
        int res=0;
        vector<string>queens(n,string(n,'.'));// build queens
        // 从第0行开始
        helper(queens,0,res);
        return res;
    }
};
