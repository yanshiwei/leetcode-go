class Solution {
    /*
    题目：
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上单元格 高于海平面的高度 。
岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
返回网格坐标 result 的 2D 列表 ，其中 result[i] = [ri, ci] 表示雨水从单元格 (ri, ci) 流动 既可流向太平洋也可流向大西洋 
    https://leetcode.cn/problems/pacific-atlantic-water-flow/solutions/754024/shui-wang-gao-chu-liu-by-xiaohu9527-xxsx/
    就是找到哪些点 可以同时到达太平洋和大西洋。 流动的方式只能从高往低流。
    从源点（格子）流向汇点（海域）是按照高度从高到低（非严格）的规则，那么反过来从海域到格子则是按照从低到高（非严格）规则进行，同时本身处于边缘的格子与海域联通。
    因此我们可以使用两遍 DFS 进行求解：分别从与当前海域直接相连的边缘格子出发，统计能够流向当前海域的格子集合，两片海域求得的集合交集即是答案。
    */
public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
        m=heights.size();
        if(m<1){
            return {};
        }
        n=heights[0].size();
        if(n<1){
            return {};
        }
        P=vector<vector<int>>(m,vector<int>(n,0));
        A=vector<vector<int>>(m,vector<int>(n,0));
        //从边缘开始DFS
        for(int i=0;i<m;i++){
            dfs(heights,P,i,0);//第一列，靠近太平洋
            dfs(heights,A,i,n-1);//最后一列，靠近大西洋
        }
        for(int j=0;j<n;j++){
            dfs(heights,P,0,j);//第一行，靠近太平洋
            dfs(heights,A,m-1,j);//最后一行，靠近大西洋
        }  
        return ans;     
    }
    void dfs(vector<vector<int>>& heights,vector<vector<int>>& visited,int i,int j){
        //出界
        if(i<0||j<0||i>=m&&j>=n){
            return;
        }
        //已经访问过
        if(visited[i][j]){
            return;
        }
        //标记
        visited[i][j]=1;
        // 每遍历完一个点后检查这个点是否能从 P 和 A 达到，见 P，A 定义
        // 如果可以则加入答案 ans
        if(P[i][j]&&A[i][j]){
            ans.push_back({i,j});
        }
        //上下左右深搜 从低到高
        if(i-1>=0&&heights[i-1][j]>=heights[i][j]){
            dfs(heights,visited,i-1,j);
        }
        if(i+1<m&&heights[i+1][j]>=heights[i][j]){
            dfs(heights,visited,i+1,j);
        }
        if(j-1>=0&&heights[i][j-1]>=heights[i][j]){
            dfs(heights,visited,i,j-1);
        }
        if(j+1<n&&heights[i][j+1]>=heights[i][j]){
            dfs(heights,visited,i,j+1);
        }      
    }
    int m;
    int n;
    // P 用于记录从太平洋出发所能达到的点
    // A 用于记录从大西洋出发所能达到的点
    vector<vector<int>> P, A, ans;
};
