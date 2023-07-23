class Solution {
    /*
    给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
    */
public:
    int numIslands(vector<vector<char>>&grid){
        int res=0;
        int m=grid.size();
        if(m<1){
            return 0;
        }
        int n=grid[0].size();
        if(n<1){
            return 0;
        }
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]=='1'){
                    res++;
                    dfs(grid,i,j);
                }
            }
        }
        return res;
    }
    void dfs(vector<vector<char>>& grid,int i,int j){
        // 岛屿问题，首先判有没有出界
        if(i<0||j<0||i>=grid.size()||j>=grid[0].size()){
            return;
        }
        // 岛屿问题，然后判断是不是岛屿
        if(grid[i][j]=='0'){
            return;
        }
        // 岛屿问题，然后判断岛屿访问过
        if(grid[i][j]=='2'){
            return;
        }        
        // 设置访问过，防止后面又访问到陷入循环
        grid[i][j]='2';
        // 岛屿问题，四个方向遍历
        dfs(grid,i+1,j);
        dfs(grid,i-1,j);
        dfs(grid,i,j-1);
        dfs(grid,i,j+1);
    }
    // int numIslands(vector<vector<char>>& grid) {
    //     // dfs https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
    //     int res=0;
    //     int m=grid.size();
    //     int n=grid[0].size();
    //     for(int i=0;i<m;i++){
    //         for(int j=0;j<n;j++){
    //             if(grid[i][j]=='1'){
    //                 //land
    //                 dfs(grid,i,j);
    //                 //用来统计有多少岛屿，岛屿是由多个陆地组成的，概念不一样
    //                 res++;
    //             }
    //         }
    //     }
    //     return res;
    // }
    // void dfs(vector<vector<char>>& grid,int i,int j){
    //     // overflow
    //     if(i<0||j<0||i>=grid.size()||j>=grid[0].size()){
    //         return;
    //     }
    //     // has visited or is sea
    //     if(grid[i][j]!='1'){
    //         return;
    //     }
    //     grid[i][j]='2';
    //     dfs(grid,i-1,j);
    //     dfs(grid,i+1,j);
    //     dfs(grid,i,j-1);
    //     dfs(grid,i,j+1);
    // }
};
