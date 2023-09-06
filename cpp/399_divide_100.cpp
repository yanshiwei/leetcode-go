class Solution {
    /*
    题目：
    你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
题意：
dfs+https://leetcode.cn/problems/evaluate-division/solutions/512094/dfsxiang-jie-pou-xi-guo-cheng-qi-shi-hen-aqin/
    */
public:
    vector<double>res;
    bool Nofond;
    // 注意这里常引用是方便dfs调用时直接临时表达式生成的常引用
    void dfs(unordered_map< string , vector<pair<string,double>> >& g, unordered_map<string,int>& visit,string source,string target, const double &factor){
        if(Nofond==false){
            // 之前已经找到，不用再处理，快速返回
            return;
        }
        if(source==target){
            Nofond=false;
            res.push_back(factor);
            return;
        }
        for(int j=0;j<g[source].size();j++){
            //检查与source相连接的点，是否已经访问过了。没访问过继续dfs
            if(visit[g[source][j].first]==0){
                // 与source相连接的点没有被访问过
                visit[g[source][j].first]=1;
                dfs(g,visit,g[source][j].first,target,factor*g[source][j].second);
                visit[g[source][j].first]=0;
            }
        }
    }
    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        unordered_map<string,vector<pair<string,double>>>g;
        unordered_map<string,int>visit;
        //构建无向图，a-b的value是3 的话 ，b-a是3的倒数
        for(int i=0;i<equations.size();i++){
            //a-b的value是3
            g[equations[i][0]].push_back(make_pair(equations[i][1],values[i]));
            //b-a的value是1.0/3
            g[equations[i][1]].push_back(make_pair(equations[i][0],1.0/values[i]));
        }
        //遍历queries，对每一组进行dfs计算结果。
        //如果相连接，就把 路径上的权值相乘就是结果
        for(int i=0;i<queries.size();i++){
            //如果query是不存在的，直接出结果：-1
            if(g.count(queries[i][0])==0||g.count(queries[i][1])==0){
                res.push_back(-1.0);
                continue;
            }
            //设置一个全局变量，如果进行dfs后，queries[0]到不了queries[1]，Nofond = true;
            Nofond=true;
            visit[queries[i][0]]=1;
            dfs(g,visit,queries[i][0],queries[i][1],1.0);
            visit[queries[i][0]]=0;
            if(Nofond){
                res.push_back(-1.0);
            }
        }
        return res;
    }
};
