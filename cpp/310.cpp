class Solution {
    /*
    树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。
请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
    题解：
    https://leetcode.cn/problems/minimum-height-trees/solutions/1398039/by-a-fei-8-hm2n/
    里面图对应示例2
    */
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        vector<int> res;
        if(n<1){
            return res;
        }
        if(n==1){
            res.push_back(0);
            return res;
        }
        // bfs+邻接表https://blog.csdn.net/qq_64691289/article/details/126478819
        vector<vector<int>>adjMaxt(n);// 模拟邻接点表,每一行对应一个节点的邻接点列表
        vector<int> degrees(n);//入度
        for(auto &e:edges){
            adjMaxt[e[0]].push_back(e[1]);
            adjMaxt[e[1]].push_back(e[0]);
            degrees[e[0]]++;
            degrees[e[1]]++;
        }
        queue<int>nodes;
        // 将入度为1的叶子节点存入队列
        // 入度为1的意义 入度为1的点基本不会作为最终答案【除了只有两个点的情况】， 
        // 因为与它相连的点（入度为1所以只有这一个点）到其他点的距离，永远比它到这些点的距离小1，
        // 以相连点为根会比入度为1的点为根最小高度更小（小于等于）。 我们刨去所有入度为1的点以后，
        // 整个图有了一个新的入度，又同样有了新的一些入度为1的点，重复上面的讨论。
        // 当前的图(初始的图或者删掉了前几层叶子节点的图)，我们要找的满足题设的根节点所在位置只有两种可能，一种在当前图的叶子节点(即度为1的节点)，一种为内部节点，若选择某叶子节点1，则该叶子节点1与任意其他叶子节点2的距离必定为叶子1-内部节点x-叶子2，深度为这三段边之和，必大于等于max(内部x-叶子1，内部x-叶子2)，所以相比于叶子节点，解空间只能出现在内部节点，每层情况都是这样，所以我们要剥开叶子节点直到再无可分的内部节点为止。
        // 故更靠近里面的节点才有可能得到最小高度
        for(int i=0;i<degrees.size();i++){
            if(degrees[i]==1){
                nodes.push(i);
            }
        }
        while(nodes.empty()==false){
            int curSize=nodes.size();
            res.clear();//若仍存在度1的 则上一轮的肯定不是最终结果
            for(int i=0;i<curSize;i++){
                auto curNode=nodes.front();
                nodes.pop();
                res.push_back(curNode);
                // 处理该节点的邻接点
                for(int j=0;j<adjMaxt[curNode].size();j++){
                    degrees[adjMaxt[curNode][j]]--;
                    if(degrees[adjMaxt[curNode][j]]==1){
                        nodes.push(adjMaxt[curNode][j]);
                    }
                }
            }
        }
        return res;
    }
};
