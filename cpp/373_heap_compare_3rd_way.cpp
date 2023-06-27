class Solution {
    /*
    题目：
    给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
    题解：
    TOP K最小，大顶堆
    队列保存<idx1,idx2>,分别表述nums1/nums2的索引；
    初始化时先把所有的 nums1 的索引入队，nums2只入第一个。
    每次弹出一个，第一个弹出的就是nums1[0] + nums2[0]（两个数组都是升序的条件，结果中最小的组合就是这一对）；接下来次小的是什么呢？ 是 nums1[0] + nums2[1] 还是 nums1[1] + nums2[0]，这时候就需要大顶堆作用，由堆完成比较并且pop。再把 index2 后移一位继续添加到优先级队列中，依次往复，最终弹出 k 次就是我们的结果。
    */
public:
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        // 由于队列保存<idx1,idx2>，没有保存nums相关信息，为了直接使用这里引入函数变量。这也是priority_queue比较函数类的第三种方式
        auto greater=[&nums1,&nums2](const pair<int,int>&a,const pair<int,int>&b){
            return nums1[a.first]+nums2[a.second]>nums1[b.first]+nums2[b.second];
        };
        int m=nums1.size();
        int n=nums2.size();
        vector<vector<int>>res;
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(greater)> maxHeap(greater);//decltype(func)转为函数类型
        // 把 nums1 的所有索引入队，nums2 的索引初始时都是 0
        // 优化：最多入队 k 个就可以了，因为提示中 k 的范围较小，这样可以提高效率
        for(int i=0;i<min(k,m);i++){
            maxHeap.push(make_pair(i,0));
        }
        while(k>0&&!maxHeap.empty()){
            auto xy=maxHeap.top();
            maxHeap.pop();
            res.push_back(vector<int>{nums1[xy.first],nums2[xy.second]});
           // 将 index2 加 1 之后继续入队
           xy.second++;
           if(xy.second<n){
               maxHeap.push(xy);
           } 
           k--;
        }
        return res;
    }
};
