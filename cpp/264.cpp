class Solution {
public:
// 类似test17.09。从小到大第n个数，可以用最小堆，第n次从最小堆取出的即可。
// 每次取出堆顶元素x，则x是当前最小值，由于2x 3x 5x均符合要求故加入堆
// 为了避免重复 加入哈希去重
    int nthUglyNumber(int n) {
        unordered_set<int>se;//insert erase size empty find count
        priority_queue<long,vector<long>,greater<long>>q;// push pop top size empty;long aoivd overflow
        // init
        se.insert(1L);
        q.push(1L);
        vector<int>lists={2,3,5};
        int res=0;
        for(int i=0;i<n;i++){
            auto cur=q.top();
            q.pop();
            res=int(cur);
            //由于2x 3x 5x均符合要求gu故加入堆
            for(int j=0;j<lists.size();j++){
                long next=cur*lists[j];
                // 如果没有重复
                if(se.count(next)==false){
                    se.insert(next);
                    q.push(next);
                }
            }
        }
        return res;
    }
};
