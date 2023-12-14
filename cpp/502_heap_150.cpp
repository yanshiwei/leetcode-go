class Solution {
    /*
    本质还是贪心，只不过多了需要满足贪心的条件而已，首先由于启动资金是固定的，
    你得从小往大去观察哪些最小资本合适，这就需要定义一个pair，最小资本+最大利润，
    在满足最小资本的条件下，找到最大的利润加到启动资金里面，这样不断循环遍历，
    直到找到k个或者遍历结束，返回最多资本即可，
    pair的排序利用sort函数，找到满足条件的最大利润利用最大堆priority_queue

    */
public:
    int findMaximizedCapital(int k, int w, vector<int>& profits, vector<int>& capital) {
        int n=profits.size();
        vector<pair<int,int>>capital_profits;
        for(int i=0;i<n;i++){
            capital_profits.push_back(pair<int,int>(capital[i],profits[i]));
        }
        // 根据最小资本排序
        sort(capital_profits.begin(),capital_profits.end());
        // 定义大堆，每次获取最大资金
        priority_queue<int>maxHeap;// base vector deque default is less
        int i=0;
        while(k>0){
            // 将所有满足启动资金的项目push
            while(i<n&&capital_profits[i].first<=w){
                maxHeap.push(capital_profits[i].second);
                i++;
            }
            if(maxHeap.empty()){
                break;
            } 
            // 获取当前启动资金下最大的利润项目
            // 且利润将被添加到总资本中
            w+=maxHeap.top();
            maxHeap.pop();
            k--;
        }
        return w;
    }
};
