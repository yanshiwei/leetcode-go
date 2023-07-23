class Solution {
    /*
    给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
    */
public:
    int islandPerimeter(vector<vector<int>>& grid) {
        //dfs https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
        // 周长包括2种，一种是与网格边界相邻的周长，我们标记为黄色线，另一种是与海洋格子相邻的周长我们标记为蓝色线。
        // 当我们的 dfs 函数因为「坐标 (i, j) 超出网格范围」返回的时候，实际上就经过了一条黄色的边；而当函数因为「当前格子是海洋格子」返回的时候，实际上就经过了一条蓝色的边。这样，我们就把岛屿的周长跟 DFS 遍历联系起来了
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
                // 因为只有一个所以找到就返回
                if(grid[i][j]==1){
                    return dfs(grid,i,j);
                }
            }
        }
        return 0;
    }
    int dfs(vector<vector<int>>& grid,int i,int j){
        // 岛屿问题，防止出界
        if(i<0||j<0||i>=grid.size()||j>=grid[0].size()){
            // 出界了就到达一个边界，周长加1
            return 1;
        }
        // 岛屿问题，如果到达海洋，退出，也是一个边界，周长+1
        if(grid[i][j]==0){
            return 1;
        }
        // 岛屿问题，判断是否访问过
        if(grid[i][j]==2){
            return 0;
        }
        // 岛屿问题，设置访问标志位
        grid[i][j]=2;
        // 岛屿问题，四个方向尝试
        return dfs(grid,i+1,j)+dfs(grid,i-1,j)+dfs(grid,i,j-1)+dfs(grid,i,j+1);
    }
};
