class Solution {
    /*
    题目：
        一个N*N的棋盘,起点是左下角,一次可以走 1-6步
    从左下角开始棋盘编号从下往上蛇形递增
                    7 8 9
                    6 5 4
                    1 2 3
    如上图(其实不是图)所示:
    起点是 1,终点是 9(N*N)  (1-9为棋盘编号)
    每格棋盘单元格内容一般是-1
    如果棋盘单元格内容不是-1,而是一个数字,
    则如果跳到这里,就要传送到指定数字编号的棋盘单元格中，横行传送称为蛇，纵向传送称为梯
    如果该位置上-1，则只能向后选择，选择的空间是当前下标+1到+6
                    -1 -1 -1
                    -1 -1 -1
                    -1 -1  8
    如图:走到第三个格的位置,就会自动被传送到8的位置
    注意:如果传送到的棋盘单元格内容也不是-1,无法进行二次传送
    题解：
    求最少需要的选择的步数，BFS
    https://leetcode.cn/problems/snakes-and-ladders/?envType=study-plan-v2&envId=top-interview-150
    */
public:
    vector<int>nums;// 用一维数组记录每个方格对应的值
    void matrix2vec(vector<vector<int>>& board){
        int n=board.size();
        int k=1;
        nums.resize(n*n+1);
        while(k<=n*n){
            int row=n-1-(k-1)/n;
            int col=((k-1)/n%2==0)?(k-1)%n:(n-1)-(k-1)%n;
            nums[k++]=board[row][col];
        }
    }
    int snakesAndLadders(vector<vector<int>>& board) {
        int n=board.size();
        matrix2vec(board);
        return onePointBfs(board);
    }
    int onePointBfs(vector<vector<int>>& board){
        int n=board.size();
        // 存储是否访问过
        unordered_map<int,bool>visited;
        queue<int>qu;//base deque list
        int step=0;//最少选择数
        qu.push(1);// 从左下角出发，编号是1
        while(!qu.empty()){
            int cur_size=qu.size();
            for(int i=0;i<cur_size;i++){
                int start=qu.front();
                qu.pop();
                if(start==n*n){
                    return step;//到达最后一个格子
                }
                // 尝试后面的6个格子
                for(int j=1;j<=6;j++){
                    int next_point=start+j;
                    if(next_point>n*n){
                        // 边界结束
                        break;
                    }
                    if(nums[next_point]!=-1){
                        // 可以传送
                        next_point=nums[next_point];
                    }
                    if(visited.count(next_point)<1){
                        // 没有访问
                        qu.push(next_point);
                        visited[next_point]=true;
                    }
                }
            }
            step++;
        }  
        return -1;      
    }
};
