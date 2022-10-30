class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int res=0;
        int m=grid.size();
        int n=grid[0].size();
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]==1){
                    int cur=dfs(grid,i,j);
                    cout<<cur<<endl;
                    res=max(res,cur);
                }
            }
        }
        return res;
    }
    int dfs(vector<vector<int>>& grid,int i,int j){
        // overflow
        if(i<0||j<0||i>=grid.size()||j>=grid[0].size()){
            return 0;
        }
        // visited or sea
        if(grid[i][j]!=1){
            return 0;
        }
        grid[i][j]=2;
        return 1+dfs(grid,i-1,j)+dfs(grid,i+1,j)+dfs(grid,i,j-1)+dfs(grid,i,j+1);
    }
};
