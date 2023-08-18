class Solution {
    /*
    题目：
    你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
    https://leetcode.cn/problems/course-schedule/solutions/250377/bao-mu-shi-ti-jie-shou-ba-shou-da-tong-tuo-bu-pai-/
    bfs
    */
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int>inDegree(numCourses);// 入度数组，课号 0 到 n - 1 作为索引，通过遍历先决条件表求出对应的初始入度
        unordered_map<int,vector<int>>hash;//key：课号 value：依赖这门课的后续课（数组）
        for(int i=0;i<prerequisites.size();i++){
            inDegree[prerequisites[i][0]]++;
            hash[prerequisites[i][1]].push_back(prerequisites[i][0]);
        }
        queue<int>qu;
        //让入度为 0 的课入列，它们是能直接选的课
        for(int i=0;i<numCourses;i++){
            if(inDegree[i]==0){
                qu.push(i);
            }
        }
        int cnt=0;//选课数目
        // 然后逐个出列，出列代表着课被选，需要减小相关课的入度。
        // 如果相关课的入度新变为 0，安排它入列、再出列……直到没有入度为 0 的课可入列
        while(qu.empty()==false){
            int cur=qu.front();
            qu.pop();
            cnt++;
            vector<int> relatedCourses = hash[cur];//获取这门课对应的后续课
            for(int i=0;i<relatedCourses.size();i++){
                //依赖它的后续课的入度-1
                inDegree[relatedCourses[i]]--;
                //如果因此减为0，入列
                if(inDegree[relatedCourses[i]]==0){
                    qu.push(relatedCourses[i]);
                }
            }
        }
        if(cnt==numCourses){
            return true;
        }
        return false;
    }
};
