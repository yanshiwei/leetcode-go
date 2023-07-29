class Solution {
    //BFS
    /*
    题意：
    给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：
路径途经的所有单元格都的值都是 0 。
路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
畅通路径的长度 是该路径途经的单元格总数。
    什么时候用BFS，什么时候用DFS？
1.如果只是要找到某一个结果是否存在，那么DFS会更高效。因为DFS会首先把一种可能的情况尝试到底，才会回溯去尝试下一种情况，只要找到一种情况，就可以返回了。但是BFS必须所有可能的情况同时尝试，在找到一种满足条件的结果的同时，也尝试了很多不必要的路径； 2.如果是要找所有可能结果中最短的，那么BFS更高效。因为DFS是一种一种的尝试，在把所有可能情况尝试完之前，无法确定哪个是最短，所以DFS必须把所有情况都找一遍，才能确定最终答案（DFS的优化就是剪枝，不剪枝很容易超时）。而BFS从一开始就是尝试所有情况，所以只要找到第一个达到的那个点，那就是最短的路径，可以直接返回了，其他情况都可以省略了，所以这种情况下，BFS更高效。

BFS解法中的visited为什么可以全局使用？
BFS是在尝试所有的可能路径，哪个最快到达终点，哪个就是最短。那么每一条路径走过的路不同，visited（也就是这条路径上走过的点）也应该不同，那么为什么visited可以全局使用呢？ 因为我们要找的是最短路径，那么如果在此之前某个点已经在visited中，也就是说有其他路径在小于或等于当前步数的情况下，到达过这个点，证明到达这个点的最短路径已经被找到。那么显然这个点没必要再尝试了，因为即便去尝试了，最终的结果也不会是最短路径了，所以直接放弃这个点即可。
本题由于是矩阵。可以省略vistied数组，把遍历了的设置为1
    */
public:
    int shortestPathBinaryMatrix(vector<vector<int>>& grid) {
        int n=grid.size();
        if(n<1){
            return -1;
        }
        int m=grid[0].size();
        if(m<1){
            return -1;
        }
        if(grid[0][0]!=0||grid[n-1][m-1]!=0){
            return -1;
        }
        if(n==1){
            // becaseu n==m
            // so source == target
            return 1;
        }
        queue<pair<int,int>>qu;
        qu.push({0,0});
        grid[0][0] = 1; // 走过的标记为1
        int step=1;
        while(!qu.empty()){
            int curSize=qu.size();
            for(int i=0;i<curSize;i++){
                pair<int,int>pp=qu.front();
                qu.pop();
                int x=pp.first;
                int y=pp.second;
                // 遍历 8 个方向
                for(int j=0;j<8;j++){
                    int xx=x+dir[j][0];
                    int yy=y+dir[j][1];
                    if(xx<0||yy<0||xx>=n||yy>=m){
                        //出界
                        continue;
                    }
                    if(xx==n-1&&yy==m-1){
                        //终点
                        return step+1;//+1加上的是终点
                    }
                    if(grid[xx][yy]==0){
                        qu.push({xx,yy});
                        grid[xx][yy]=1;//标志访问过
                    }else{
                        //访问过或者原本是1
                        continue;
                    }
                }
            }
            step++;
        }
        return -1;
    }
    int dir[8][2]={{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,0},{1,1},{1,-1}};
};
