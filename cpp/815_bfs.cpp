class Solution {
    //https://leetcode.cn/problems/bus-routes/solutions/798032/java815gong-jiao-lu-xian-yi-li-jie-ba-ch-0kfv/
    // 知道source，target，求最少乘坐的公交车数量，BFS。一个层级就是一辆公交车。
public:
    int numBusesToDestination(vector<vector<int>>& routes, int source, int target) {
        if(routes.size()<1){
            return -1;
        }
        if(source==target){
            return 0;
        }
        unordered_map<int,vector<int>>s2b;//<station,{bus}>-每个站都被哪些公交车经过
        for(int i=0;i<routes.size();i++){
            for(int j=0;j<routes[i].size();j++){
                if(s2b.count(routes[i][j])==0){
                    vector<int>tmp;
                    tmp.push_back(i);
                    s2b[routes[i][j]]=tmp;
                }else{
                    s2b[routes[i][j]].push_back(i);
                }
            }
        }
        vector<bool>visited(routes.size(),false);//记录已经坐了哪些公交车,保证每个路线只乘坐过1次。
        queue<int>qu;
        qu.push(source);//从 source 车站出发（初始时不在公交车上），故也不需要设置visited[0]=true
        int step=0;//坐过多少公交车，这时候还没有上车（初始时不在公交车上）
        while(!qu.empty()){
            int curSize=qu.size();
            // 处理非换乘单一公交车情况
            for(int i=0;i<curSize;i++){
                int curStation=qu.front();
                qu.pop();
                //经过curStation的所有车
                for(int car:s2b[curStation]){
                    // 已经坐过则跳过
                    if(visited[car]){
                        continue;
                    }
                    // 之前没有坐过，继续处理
                    // 标记已经访问过的该公交车
                    visited[car]=true;
                    // 当前公交车对应所有站点
                    for(int station:routes[car]){
                        if(station==target){
                            // 找到直接返回
                            return step+1;//
                        }
                        if(station==curStation){
                            continue;
                        }
                        qu.push(station);
                    }
                }
            }
            step++;
        }
        return -1;
    }
};
