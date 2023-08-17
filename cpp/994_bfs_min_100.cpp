class Solution {
public:
    // 遍历，最短问题，BFS, see https://leetcode.cn/problems/01-matrix/
    int orangesRotting(vector<vector<int>>& grid) {
        int res=0;
        int m=grid.size();
        if(m<1){
            return 0;
        }
        int n=grid[0].size();
        if(n<1){
            return 0;
        }
        //首先将所有的 2 都入队，并且统计新鲜个数
        queue<pair<int,int>>qu;
        int fresh=0;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]==2){
                    qu.push({i,j});
                }else if(grid[i][j]==1){
                    fresh++;
                }
            }
        }
        int dir[4][2]={{-1,0},{1,0},{0,-1},{0,1}};
        while(!qu.empty()){
            int cur_size=qu.size();
            bool fresh_is_bad =false;//本轮有没有破坏新鲜水果
            for(int i=0;i<cur_size;i++){
                pair<int,int>pp=qu.front();
                qu.pop();
                int x=pp.first;
                int y=pp.second;
                // 四个方向
                for(int j=0;j<4;j++){
                    int xx=x+dir[j][0];
                    int yy=y+dir[j][1];
                    // 出界判断
                    if(xx<0||yy<0||xx>=m||yy>=n){
                        continue;
                    }
                    // 是不是新鲜
                    if(grid[xx][yy]!=1){
                        // 可能已经访问过或者是空气
                        continue;
                    }
                    // 遇到新鲜的
                    grid[xx][yy]=2;
                    fresh--;
                    qu.push({xx,yy});
                    fresh_is_bad=true;
                }
            }
            if(fresh_is_bad){
                // 本轮破坏了新鲜，每腐烂一次时间加 1
                res++;
            }
        }
        return fresh>0?-1:res;
    }
};
