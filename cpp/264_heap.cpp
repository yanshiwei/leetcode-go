class Solution {
    /*
    题目：
    给你一个整数 n ，请你找出并返回第 n 个 丑数 。
丑数 就是只包含质因数 2、3 和/或 5 的正整数。
    题解：
    从小到大第n个数，可以用最小堆，第n次从堆顶取出的元素就是第n个
    每次取出堆顶元素x，由于由于2x 3x 5x均符合要求故加入堆，此外有可能有重复所以需要
    加入去重逻辑
    */
public:
    int nthUglyNumber(int n) {
        unordered_set<long>hash;
        priority_queue<long,vector<long>,greater<long>>minHeap;
        hash.insert(1);
        minHeap.push(1);
        vector<int>factors={2,3,5};
        int res=0;
        for(int i=0;i<n;i++){
            // 每次选出最小值
            long cur=minHeap.top();
            minHeap.pop();
            res=int(cur);
            for(auto&fa:factors){
                long next=cur*fa;
                if(hash.count(next)==0){
                    hash.insert(next);
                    minHeap.push(next);
                }
            }
        }
        return res;
    }
};
