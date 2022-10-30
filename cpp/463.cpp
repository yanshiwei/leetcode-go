class Solution {
    /*
    给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
    */
public:
    int islandPerimeter(vector<vector<int>>& grid) {
        //dfs https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
        // 贡献的周长部分就是4条边中靠近grid边缘或0的边的数量 全部加在一起就是周长了
        for(int i=0;i<grid.size();i++){
            for(int j=0;j<grid[0].size();j++){
                if(grid[i][j]==1){
                    // 题目限制只有一个岛屿，计算一个即可
                    return dfs(grid,i,j);
                }
            }
        }
        return 0;
    }
    int dfs(vector<vector<int>>& grid,int i,int j){
        //与边界相邻，周长加1
        if(i<0||j<0||i>=grid.size()|j>=grid[0].size()){
            return 1;
        }
        //与水域相邻，边界加1
        if(grid[i][j]==0){
            return 1;
        }
        if(grid[i][j]!=1){
            return 0;
        }
        grid[i][j]=2;
        return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1);
    }
};
