class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        if(nums.size()<k||k<1){
            return {};
        }
        priority_queue<pair<int,int>>q;//默认less，大顶堆.first is valu, second is index
        // 为了方便判断堆顶元素与滑动窗口的位置关系，我们可以在优先队列中存储二元组<value,index>
        vector<int>res;
        // 初始时，我们将数组 nums的前 k个元素放入优先队列中
        for(int i=0;i<k;i++) {
            q.push(make_pair(nums[i],i));
        }
        res.push_back(q.top().first);
        for(int i=k;i<nums.size();i++) {
            q.push(make_pair(nums[i],i));
            // 堆顶元素就是最大值，但是可能已经不在窗口内
            // 这个值在数组 nums中的位置出现在滑动窗口左边界的左侧
            while(q.top().second<=i-k){
                q.pop();
            } 
            res.push_back(q.top().first);
        }
        return res;
    }
};
