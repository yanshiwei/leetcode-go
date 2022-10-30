class Solution {
    /*
    给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
返回执行此操作后，grid 中最大的岛屿面积是多少？
岛屿 由一组上、下、左、右四个方向相连的 1 形成
    */
public:
    unordered_map<int,int>lnd;//岛屿编号 -> 岛屿面积的 map
    int largestIsland(vector<vector<int>>& grid) {
        lnd[0]=0;
        int value=2;// 岛屿编号,因为后面要标柱grid[i][j]而grid[i][j]0 1已经有含义故从2开始
        // 第一次DFS 遍历陆地格子，计算每个岛屿的面积并标记岛屿
        for(int i=0;i<grid.size();i++){
            for(int j=0;j<grid[0].size();j++){
                if(grid[i][j]==1){
                    int cur=dfs(grid,i,j,value);
                    lnd[value]=cur;
                    value++;
                }
            }
        }
        // 第二遍 DFS 遍历海洋格子，观察每个海洋格子相邻的陆地格子
        int res=0;
        for(int i=0;i<grid.size();i++){
            for(int j=0;j<grid[0].size();j++){
                // 依次尝试填海
                int cur=fillSeaDfs(grid,i,j);
                res=max(res,cur);
            }
        }
        return res;
    }
    int fillSeaDfs(vector<vector<int>>& grid,int i,int j){
        if(grid[i][j]!=0){
            // not sea 返回该岛屿面积
            return lnd[grid[i][j]];
        }
        int res=0;
        unordered_set<int>adjs;// 去重复的岛屿下标
        // 四个方向岛屿
        if(inArea(grid,i-1,j)&&grid[i-1][j]>0){
            adjs.insert(grid[i-1][j]);
        }
        if(inArea(grid,i+1,j)&&grid[i+1][j]>0){
            adjs.insert(grid[i+1][j]);
        }
        if(inArea(grid,i,j-1)&&grid[i][j-1]>0){
            adjs.insert(grid[i][j-1]);
        }
        if(inArea(grid,i,j+1)&&grid[i][j+1]>0){
            adjs.insert(grid[i][j+1]);
        }
        // 四个方向去掉本身后的面积之和
        for(int adj:adjs){
            res+=lnd[adj];
        }
        return res+1;// 加上sea一个
    }
    int dfs(vector<vector<int>>& grid,int i,int j,int value){
        // overflow
        if(!inArea(grid,i,j)){
            return 0;
        }
        if(grid[i][j]!=1){
            return 0;
        }
        grid[i][j]=value;
        return 1+dfs(grid,i-1,j,value)+dfs(grid,i+1,j,value)+dfs(grid,i,j-1,value)+dfs(grid,i,j+1,value);
    }
    bool inArea(vector<vector<int>>& grid, int r, int c) {
        return 0 <= r && r < grid.size()
            && 0 <= c && c < grid[0].size();
    }
};
