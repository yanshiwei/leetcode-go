class Solution {
    class UinonFind{
        public:
            // 初始化，n个结点，每个单独就是一个集合
            UinonFind(int n){
                father.resize(n);
                for(int i=0;i<n;i++){
                    father[i]=i;// 按照定义，根结点保存自身引用
                }
            }
            int find(int x){
                int origin=x;
                while(x!=father[x]){
                    x=father[x];
                }
                father[origin]=x;//路径压缩,origin的代表结点也是a
                return x;
            }
            void merge(int x,int y){
                int father_x=find(x);
                int father_y=find(y);
                if(father_x==father_y){
                    return;
                }
                father[father_x]=father_y;//x所在集合合入y所在集合
            }
        private:
            vector<int>father;
    };
public:
    int findCircleNum(vector<vector<int>>& isConnected) {
        int n=isConnected.size();
        UinonFind uf(n);
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(isConnected[i][j]){
                    uf.merge(i,j);
                }
            }
        }
        // 此时已经并查集已经把边链接好了 
        unordered_set<int>uset;
        for(int i=0;i<n;i++){
            uset.insert(uf.find(i));
        }
        return uset.size();
    }
};
