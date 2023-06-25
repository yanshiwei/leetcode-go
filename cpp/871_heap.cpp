class Solution {
    /*
    题目：
    汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。
沿途有加油站，用数组 stations 表示。其中 stations[i] = [positioni, fueli] 表示第 i 个加油站位于出发位置东面 positioni 英里处，并且有 fueli 升汽油。
假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。它每行驶 1 英里就会用掉 1 升汽油。当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。
为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。
注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。
    题解：
    先不考虑还剩多少油，考虑当前油量下，最多可以走多远。
    具体来说，当走到一个加油站时，先不加油，先存储用的时候再拿出来。
    什么时候要用呢？就是无法到达最近的加油站时，就需要用， 为了保证加油次数最少，首先拿出的油应该是最多的，故这里引入引入了大堆，每次pop都是pop出序列的最大值
    */
public:
    int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
        if(startFuel>=target){
            // 初始油就足以到达，不用加油
            return 0;
        }
        priority_queue<int>maxHeap;//默认大堆
        int cnt=0;
        int curFuel=startFuel;
        for(int i=0;i<stations.size();i++){
            // 判断当前油量能否到达最近的加油站
            while(curFuel<stations[i][0]){
                // 当前油量不能到达最近的加油站， 拿出存储的最大加油储备
                if(maxHeap.empty()){
                    return -1;
                }
                curFuel+=maxHeap.top();//pop出当前最大加油量
                maxHeap.pop();
                cnt++;
            }
            // 当前油量无需加油（或者加油后）就可以到最近点，先不加油存储到堆
            maxHeap.push(stations[i][1]);
        }
        // 过完所有加油站后，判断当前油量能否到达target
        while(curFuel<target){
            // 当前油量不能到达最近的target， 拿出存储的最大加油储备
            if(maxHeap.empty()){
                return -1;
            }
            curFuel+=maxHeap.top();//pop出当前最大加油量
            maxHeap.pop();
            cnt++;
        }
        return cnt;
    }
};
