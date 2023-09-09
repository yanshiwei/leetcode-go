class Solution {
    /*
    https://leetcode.cn/problems/game-of-life/solutions/139404/ju-zhen-wen-ti-tong-yong-jie-fa-by-freshrookie/?envType=study-plan-v2&envId=top-interview-150
   属于01，不用更新，向右移动后死亡；
    */
public:
    vector<int>dx={0,0,-1,1,1,-1,1,-1};
    vector<int>dy={-1,1,0,0,-1,1,1,-1};
    int cntFor8Dirs(vector<vector<int>>& board,int x,int y,int m,int n){
        int res=0;
        for(int k=0;k<8;k++){
            int curx=x+dx[k];
            int cury=y+dy[k];
            if(curx<0||curx>=m||cury<0||cury>=n){
                continue;
            }
            //如果该位置为0，当前轮是死的，不用考虑
            //如果该位置为1，当前轮是活的，算进去
            //如果该位置为2，即10状态，表示当前轮是死的，下一轮才是活，不需要算进去
            //度过该位置为3，即11状态，表示当前轮是活得，下一轮是活得，算尽
            res+=(board[curx][cury]&1);
        }
        return res;
    }
    void gameOfLife(vector<vector<int>>& board) {
        int m=board.size();
        if(m<1){
            return;
        }
        int n=board[0].size();
        if(n<1){
            return;
        }
        // 第一次遍历，把需要变化状态的位置保存为中间值
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                int curx;
                int cury;
                int live=0;
                for(int k=0;k<8;k++){
                    curx=i+dx[k];
                    cury=j+dy[k];
                    if(curx<0||curx>=m||cury<0||cury>=n){
                        continue;
                    }
                    // 这里的0,1是题目里合理的值，0到0不用改，1到1也不用更改，然后如果0要变1，我们用中间值-1记录，如果1要变0
                    // ，我们用中间值2来记录。
                    if(board[curx][cury]==1||board[curx][cury]==2){
                        live++;
                    }
                }
                if(board[i][j]==0){
                    //原始0状态的，下一个状态要么是0要么是1
                    //经过处理后，变成0或者1，还是0的的不处理了仍然照旧就行；
                    //只处理0-->1的设置board[i][j]=-1，后续再遇到board[i][j]==-1就不重复处理
                    if(live==3){
                        board[i][j]=-1;
                    }
                }else{
                    //原始1状态的，下一个状态要么是1要么是0
                    //经过处理后，变成1或者0，还是1的的不处理了仍然照旧就行；
                    //只处理1-->0的设置board[i][j]=2，后续再遇到board[i][j]==2就不重复处理
                    if(live<2||live>3){
                        board[i][j]=2;
                    }
                }
            }
        }
        // 遍历更新为最终态
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(board[i][j]==2){
                    // 1-->0
                    board[i][j]=0;
                }else if(board[i][j]==-1){
                    // 0--->1
                    board[i][j]=1;
                }
            }
        }
    }
};
