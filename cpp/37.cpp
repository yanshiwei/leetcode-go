class Solution {
    /*
    题目：
    数独的解法需 遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。
    题解：
    解数独思路：
    类似人的思考方式去尝试，行，列，还有 3*3 的方格内数字是 1~9 不能重复。
    我们尝试填充，如果发现重复了，那么擦除重新进行新一轮的尝试，直到把整个数组填充完成。
    算法步骤:
    1 数独首先行，列，还有 3*3 的方格内数字是 1~9 不能重复。
    2 声明布尔数组，表明行列中某个数字是否被使用了， 被用过视为 true，没用过为 false。初始化布尔数组，表明哪些数字已经被使用过了。
    3 尝试去填充数组，只要行，列， 还有 3*3 的方格内 出现已经被使用过的数字，我们就不填充，否则尝试填充。
    4 如果填充失败，那么我们需要回溯。将原来尝试填充的地方改回来。
    5 递归直到数独被填充完成。
    */
public:
// 三个布尔数组 表明 行, 列, 还有 3*3 的方格的数字是否被使用过
    vector<vector<bool>>rowUsed;
    vector<vector<bool>>colUsed;//二维数组存储的是[列下标:0-8][数独数字:1-9]，第二维度上10为了写代码方便, 0是没用
    vector<vector<vector<bool>>>boxUsed;
    bool solveSudokuHelper(vector<vector<char>>& board,vector<vector<bool>>&rowUsed,vector<vector<bool>>&colUsed,vector<vector<vector<bool>>>&boxUsed,int row,int col){
        // 边界校验, 如果已经填充完成, 返回true, 表示一切结束
        if(col==board[0].size()){
            //由于是按照列递归，故如果到边界要重置，行加1
            col=0;
            row++;
            if(row==board.size()){
                return true;
            }
        }
        // 是空则尝试填充, 否则跳过继续尝试填充下一个位置
        if(board[row][col]=='.'){
            // 尝试填充1~9
            for(int num=1;num<=9;num++){
                bool canPut=!(rowUsed[row][num]||colUsed[col][num]||boxUsed[row/3][col/3][num]);
                if(canPut){
                    rowUsed[row][num]=true;
                    colUsed[col][num]=true;
                    boxUsed[row/3][col/3][num]=true;
                    board[row][col]=char(num+'0');
                    if(solveSudokuHelper(board,rowUsed,colUsed,boxUsed,row,col+1)){
                        return true;
                    }
                    //如果填充失败，那么我们需要回溯。将原来尝试填充的地方改回来
                    rowUsed[row][num]=false;
                    colUsed[col][num]=false;
                    boxUsed[row/3][col/3][num]=false;
                    board[row][col]='.';                    
                }
            }
        }else{
            return solveSudokuHelper(board,rowUsed,colUsed,boxUsed,row,col+1);//按照列递归
        }
        return false;
    }
    void solveSudoku(vector<vector<char>>& board) {
        // init 
        rowUsed=vector<vector<bool>>(9,vector<bool>(10,false));
        colUsed=vector<vector<bool>>(9,vector<bool>(10,false));
        boxUsed=vector<vector<vector<bool>>>(3,vector<vector<bool>>(3,vector<bool>(10,false)));
        for(int i=0;i<board.size();i++){
            for(int j=0;j<board[0].size();j++){
                int num=board[i][j]-'0';
                if(num>=1&&num<=9){
                    // 已经放置了数字num
                    rowUsed[i][num]=true;
                    colUsed[j][num]=true;
                    boxUsed[i/3][j/3][num]=true;
                }
            }
        }
        // 递归尝试填充数组 
        solveSudokuHelper(board,rowUsed,colUsed,boxUsed,0,0);
    }
};
