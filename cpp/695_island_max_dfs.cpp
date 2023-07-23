class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int m=grid.size();
        if(m<1){
            return 0;
        }
        int n=grid[0].size();
        if(n<1){
            return 0;
        }
        int res=0;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]==1){
                    int curRes=dfs(grid,i,j);
                    res=max(res,curRes);
                }
            }
        }
        return res;
    }
    int dfs(vector<vector<int>>& grid,int i,int j){
        // 岛屿问题，首先判断是否出界
        if(i<0||j<0||i>=grid.size()||j>=grid[0].size()){
            return 0;
        }
        // 岛屿问题，判断是不是岛屿
        if(grid[i][j]==0){
            return 0;
        }
        // 岛屿问题，判断是否访问过
        if(grid[i][j]==2){
            return 0;
        }
        // 岛屿问题，设置访问过
        grid[i][j]=2;
        // 岛屿问题，四个方向尝试
        return 1+dfs(grid,i+1,j)+dfs(grid,i-1,j)+dfs(grid,i,j-1)+dfs(grid,i,j+1);
    }
};
