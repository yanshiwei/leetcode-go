class Solution {
    /*
    给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
示例 1:
输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
     题解：
     https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/solutions/1210221/bu-chong-guan-fang-ti-jie-you-xian-dui-l-htf8/?envType=study-plan-v2&envId=top-interview-150
     TOP K最小，大顶堆
     队列保存pair<idx1,idx2>,分别表述nums1/nums2的索引；
     1.初始化时先把所有的 nums1 的索引入队，nums2只入第一个。
     2.每次弹出一个，第一个弹出的就是nums1[0] + nums2[0]（两个数组都是升序的条件，结果中最小的组合就是这一对）；
     3.接下来次小的是什么呢？ 是 nums1[0] + nums2[1] 还是 nums1[1] + nums2[0]，这时候就需要大顶堆作用，由堆完成比较并且pop。
     4.再把 index2 后移一位继续添加到优先级队列中，依次往复，最终弹出 k 次就是我们的结果。
    堆自定义比较的三种方案：
    1 类内重载操作符
    bool operator<(const node&, const node&b){
        return a.val < b.val;// 按照value从大到小排列
    }
    priority_queue<node>qu;
    2 自定义仿函数
    struct cmp{
        bool operator()(const node&, const node&b){
            return a.val >b.val;// 按照value从小到大排列
        }
    }
    priority_queue<node,vector<node>,cmp>qu;
    3 类内定义友元操作类重载函数
    struct node{
        int val;
        friend bool operator(const node&, const node&b){
            return a.val < b.val;// 按照value从大到小排列
        }
    }
    priority_queue<node>qu;
    */
public:
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        // 由于队列保存<idx1,idx2>，没有保存nums相关信息，为了直接使用这里引入函数变量。这也是priority_queue比较函数类的第四种方式
        auto greater=[&nums1,&nums2](const pair<int,int>&a,const pair<int,int>&b){
            return nums1[a.first]+nums2[a.second]>nums1[b.first]+nums2[b.second];
        };
        int m=nums1.size();
        int n=nums2.size();
        vector<vector<int>>res;
        priority_queue<pair<int,int>,vector<pair<int,int>>,decltype(greater)> maxHeap(greater);//base vector deque
        // 把 nums1 的所有索引入队，nums2 的索引初始时都是 0
        // 优化：最多入队 k 个就可以了，因为提示中 k 的范围较小，这样可以提高效率
        for(int i=0;i<min(m,k);i++){
            maxHeap.push(pair<int,int>(i,0));//or maxHeap.push(make_pair(i,0));
        }
        while(k>0&&!maxHeap.empty()){
            pair<int,int>xy=maxHeap.top();
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
