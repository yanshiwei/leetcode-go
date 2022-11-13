class Solution {
    /*
    题意：
    你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1 。
    题解：
    https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/solutions/232460/wang-ge-zhong-de-zui-duan-lu-jing-bfssuan-fa-shi-x/

    */
public:
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
