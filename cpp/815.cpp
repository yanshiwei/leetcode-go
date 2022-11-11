class Solution {
    //https://leetcode.cn/problems/bus-routes/solutions/798032/java815gong-jiao-lu-xian-yi-li-jie-ba-ch-0kfv/
public:
    int numBusesToDestination(vector<vector<int>>& routes, int source, int target) {
        if(routes.size()<1){
            return -1;
        }
        if (source==target){
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
        //记录已经坐了哪些公交车
        vector<int>memery2b(routes.size(),0);
        // bfs-收集当前station辐射到的station(指的是不换乘情况下)如实例1中1可以到达2 7
        // 一个层级就是一辆车
        queue<int>qu;
        qu.push(source);
        int count=1;//坐过多少公交车
        while(qu.empty()==false){
            int sz=qu.size();
            // 处理非换乘单一公交车情况
            for(int i=0;i<sz;i++){
                int curStation=qu.front();
                qu.pop();
                //经过cur的所有车
                for(int car:s2b[curStation]){
                    if(memery2b[car]==1){
                        // 已经坐过
                        continue;
                    }
                    memery2b[car]=1;  //标记已经访问过的car
                    // 该车经过的站点
                    for(int s:routes[car]){
                        if(s==target){
                            return count;//最短
                        }
                        if(s==curStation){
                            continue;
                        }
                        qu.push(s);
                    }
                }
            }
            cout<<endl;
            // 换乘下一个公交车
            count++;
        }
        return -1;
    }
};
