class Solution {
    /*
    题目：
    现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
    题解：
    [1,3],[1,4]表明3/4课程没有先修课程，直接可以上完，而1依赖3和4，只有3和4上完才能流向1
    故1的入度为1，3和4入度为0。
    拓扑排序，类似207题。
    */
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int>inDegrees(numCourses);//inDegrees[i]代表课程i的入度，也就是i所依赖的先修课程数
        unordered_map<int,vector<int>>table;//key is课程，value是依赖该课程的所有后续课程
        for(int i=0;i<prerequisites.size();i++){
            inDegrees[prerequisites[i][0]]++;
            table[prerequisites[i][1]].push_back(prerequisites[i][0]);
        }
        // 先下线没有依赖的课程
        queue<int>qu;
        for(int i=0;i<inDegrees.size();i++){
            if(inDegrees[i]==0){
                qu.push(i);
            }
        }
        // 拓扑排序
        vector<int>res;
        while(!qu.empty()){
            int cur=qu.front();
            qu.pop();
            res.push_back(cur);
            // 下线后减去依赖该课程后续课程的一个度
            vector<int>afterCourses=table[cur];
            for(int i=0;i<afterCourses.size();i++){
                inDegrees[afterCourses[i]]--;
                if(inDegrees[afterCourses[i]]==0){
                    qu.push(afterCourses[i]);
                }
            }
        }
        return res.size()==numCourses?res:vector<int>();
    }
};
