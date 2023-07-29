class Solution {
    /*
    题意：
    你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1 。
    题解：
    https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/solutions/232460/wang-ge-zhong-de-zui-duan-lu-jing-bfssuan-fa-shi-x/

    */
public:
    int shortestPath(vector<vector<int>>& grid, int k){
        // 图片解释：https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/solutions/1740488/tu-jie-by-zhug-4-cst1/ 
        // BFS搜索节点访问标识, 此题要求有k个消除障碍的机会，所以每个节点除了标记是否被访问过
        // 还要记录搜索到此节点时消除了几个障碍。
        int m=grid.size();
        if(m<1){
            return -1;
        }
        int n=grid[0].size();
        if(n<1){
            return -1;
        }
        if(m==1&&n==1){
            //只有一个格子，且题目要求grid[0][0] == grid[m-1][n-1] == 0
            return 0;
        }
        // visited:维度为[n, m, n*m+1]，visited[i][j][k]表示原始矩阵grid中第i行第j列，移除了k个障碍的情况下，grid[i][j]是否被访问过
        // BFS搜索节点访问标识, 此题要求有k个消除障碍的机会，所以每个节点除了标记是否被访问过
        // 还要记录搜索到此节点时消除了几个障碍。消除相同障碍的下一层节点 可以剪枝（因为有相同代价更早的节点了）
        // 例子：k=1, BFS是按层级来的，绕道的层级扩展越多
        // 坐标(0,2)可以为消除(0,1)障碍过来的 visited[0][2][1] = 1
        // 也可能为不消除任何障碍过来的 visited[0][2][0] = 1
        vector<vector<vector<int>>>visited(m,vector<vector<int>>(n,vector<int>(k+1,0)));
        queue<Point> qu;
        // 初始时，左上角节点进队，此时移除的障碍数为0
        qu.push(Point(0,0,0));
        visited[0][0][0]=1;
        int step=1;//起始点算一步
        while(!qu.empty()){
            int curSize=qu.size();
            // BFS搜索-遍历相同层级下所有节点
            for(int i=0;i<curSize;i++){
                // 取出当前状态更改的障碍数以及坐标位置
                Point pp=qu.front();
                qu.pop();
                int x=pp.x;
                int y=pp.y;
                int cnt=pp.cnt;
                // 对当前节点四个方向进行识别处理
                for(int j=0;j<4;j++){
                    int xx=x+dir[j][0];
                    int yy=y+dir[j][1];
                    //出界
                    if(xx<0||yy<0|xx>=m||yy>=n){
                        continue;
                    }
                    // 穿越障碍次数已满
                    if(grid[xx][yy]==1&&cnt>=k){
                        continue;
                    }
                    // 搜索到目标节点则直接返回结果
                    if(xx==m-1&&yy==n-1){
                        return step;
                    }
                    int nextCnt=grid[xx][yy]==1?cnt+1:cnt;
                    // 四个方向节点是否被访问过（第三维度）
                    if(visited[xx][yy][nextCnt]==1){
                        //消除相同障碍的下一层节点 可以剪枝（因为有相同代价更早的节点了）
                        //也就是说有其他路径在小于或等于当前步数的情况下，到达过这个点，证明到达这个点的最短路径已经被找到。那么显然这个点没必要再尝试了，因为即便去尝试了，最终的结果也不会是最短路径了，所以直接放弃这个点即可。

                        continue;
                    }else{
                        visited[xx][yy][nextCnt]=1;
                        qu.push(Point(xx,yy,nextCnt));
                    }
                }
            }
            step++;
        }
        return -1;
    }
    /*
    int shortestPath(vector<vector<int>>& grid, int k) {
        int m=grid.size();
        int n=grid[0].size();
        // grid[0][0] == grid[m-1][n-1] == 0
        if(m==1&&n==1){
            return 0;
        }
        // BFS搜索节点访问标识, 此题要求有k个消除障碍的机会，所以每个节点除了标记是否被访问过
        // 还要记录搜索到此节点时消除了几个障碍。消除相同障碍的下一层节点 可以剪枝（因为有相同代价更早的节点了）
        // 例子：k=1, BFS是按层级来的，绕道的层级扩展越多
        // 坐标(0,2)可以为消除(0,1)障碍过来的 visited[0][2][1] = 1，搜索层级为2
        // 也可能为不消除任何障碍过来的 visited[0][2][0] = 1，层级为6，为扩展搜索不通障碍消除数提供区分
        // 0 1 0 0 0 1 0 0
        // 0 1 0 1 0 1 0 1
        // 0 0 0 1 0 0 1 0
        // 二维标记位置，第三维度标记 到此节点的路径处理障碍总个数
        int ***visited=new int**[m];
        for(int i=0;i<m;i++){
            visited[i]=new int *[n];
            for(int j=0;j<n;j++){
                visited[i][j]=new int[k+1];
            }
        }
        int steps=0;
        // 初始位置标记已访问
        visited[0][0][0]=1;
        queue<Point> qu;
        Point start=Point(0,0,0);
        qu.push(start);
        while(qu.empty()==false){
            steps++;
            // BFS搜索-遍历相同层级下所有节点
            int sz=qu.size();
            for(int i=0;i<sz;i++){
                auto cur=qu.front();
                qu.pop();
                int x=cur.x;
                int y=cur.y;
                int cnt=cur.cnt;
                // 对当前节点四个方向进行识别处理
                for(int j=0;j<4;j++){
                    int xx=x+dir[j][0];
                    int yy=y+dir[j][1];
                    // 越界
                    if(xx<0||yy<0||xx>=m||yy>=n){
                        continue;
                    }
                    // 穿越障碍次数已满
                    if(grid[xx][yy]==1&&cnt>=k){
                        continue;
                    }
                    // 搜索到目标节点则直接返回结果
                    if(xx==m-1&&yy==n-1){
                        return steps;
                    }
                    int nextCnt = grid[xx][yy]==1?cnt+1:cnt;//下一层节点已经使用障碍物次数
                    // 四个方向节点是否被访问过（第三维度）
                    if(visited[xx][yy][nextCnt]==1){
                        //消除相同障碍的下一层节点 可以剪枝（因为有相同代价更早的节点了）
                        //也就是说有其他路径在小于或等于当前步数的情况下，到达过这个点，证明到达这个点的最短路径已经被找到。那么显然这个点没必要再尝试了，因为即便去尝试了，最终的结果也不会是最短路径了，所以直接放弃这个点即可。
                        continue;
                    }else{
                        // 未被访问过且可以走的节点标记为访问过，对下一步节点确认状态非常重要
                        // 将下一层级节点入队列标记为已访问，可以剪枝更多节点，节省计算耗时
                        visited[xx][yy][nextCnt]=1;
                        auto newPoint=Point(xx,yy,nextCnt);
                        qu.push(newPoint);
                    }
                }
            }
        }
        return -1;
    }
    */
    int dir[4][2]={{-1,0},{1,0},{0,-1},{0,1}};
    class Point{
        public:
            int x;
            int y;
            int cnt;//已经使用消除障碍物的次数
            Point(int xx,int xy,int xcnt){
                 x=xx;
                 y=xy;
                 cnt=xcnt;
            }
    };
};
