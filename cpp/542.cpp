class Solution {
    /*
    题意：
    给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
两个相邻元素间的距离为 1 
    思路：
对于 「Tree 的 BFS」 （典型的「单源 BFS」） 大家都已经轻车熟路了：
首先把 root 节点入队，再一层一层无脑遍历就行了。
对于 「图 的 BFS」 （「多源 BFS」） 做法其实也是一样滴～，与 「Tree 的 BFS」的区别注意以下两条就 ok 辣～
Tree 只有 1 个 root，而图可以有多个源点，所以首先需要把多个源点都入队；
Tree 是有向的因此不需要标识是否访问过，而对于无向图来说，必须得标志是否访问过哦！并且为了防止某个节点多次入队，需要在其入队之前就将其设置成已访问！【 看见很多人说自己的 BFS 超时了，坑就在这里哈哈哈
做法：
根据上述思路，本题怎么做就很简单了：
首先把每个源点 000 入队，然后从各个 000 同时开始一圈一圈的向 111 扩散（每个 111 都是被离它最近的 000 扩散到的 ），扩散的时候可以设置 int[][] dist 来记录距离（即扩散的层次）并同时标志是否访问过。对于本题是可以直接修改原数组 int[][] matrix 来记录距离和标志是否访问的，这里要注意先把 matrix 数组中 1 的位置设置成 -1 （设成Integer.MAX_VALUE啦，m * n啦，10000啦都行，只要是个无效的距离值来标志这个位置的 1 没有被访问过就行辣~）
    */
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        // 多源 BFS
        //首先将所有的 0 都入队，并且将 1 的位置设置成 -1，表示该位置是 未被访问过的 
        int m=mat.size();
        int n=mat[0].size();
        queue<pair<int,int>>qu;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(mat[i][j]==0){
                    qu.push({i,j});
                }else{
                    mat[i][j]=-1;
                }
            }
        }
        int dir[4][2]={{-1,0},{1,0},{0,-1},{0,1}};
        while(qu.empty()==false){
            int sz=qu.size();
            // 遍历第一层
            for(int i=0;i<sz;i++){
                pair<int,int>cur=qu.front();
                qu.pop();
                int x=cur.first;
                int y=cur.second;
                // 四个方向
                for(int j=0;j<4;j++){
                    int xx=x+dir[j][0];
                    int yy=y+dir[j][1];
                    if(xx>=0&&yy>=0&&xx<m&&yy<n){
                        // 不越界
                        if(mat[xx][yy]==-1){
                            //未被访问过的 
                            mat[xx][yy]=mat[x][y]+1;//上一层该位置到0的距离+1格
                            qu.push({xx,yy});
                        }
                    }
                }
            }
        }
        return mat;
    }
};
