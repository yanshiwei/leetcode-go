class Solution {
    /*
    题目：
    给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
    题解：
    类似973题，间接求TOPK。
    最大TOPK，维护大小k的最小堆，当堆顶元素比当前元素小时，替换；当堆顶元素比当前元素大时，跳过，结束后堆里存的就是最大TOPK，堆顶最小
    */
public:
    // priority_queue第三个参数Compare：两个参数且返回值为bool类型的双参判断式
    // priority_queue自定义排序方法1，仿函数类
    struct greater{
        bool operator()(pair<int,int>m,pair<int,int>n){
            return m.second>n.second;
        }
    };
    // priority_queue自定义排序方法2，自定义类里面重载>
    //struct Node {
    //int size;
    //int price;
    // 重载>运算符
    //bool operator>(const Node &b) const {
    //    return this->size == b.size ? this->price < b.price : this->size > b.size;
    //}
//};
    vector<int> topKFrequent(vector<int>& nums, int k) {
        if(k<1){
            return {};
        }
        priority_queue<pair<int,int>,vector<pair<int,int>>,greater>minHeap;
        // 默认大堆顶less<int>,对于pair是先比较first，相同则比较second
        // 故需要重新定义比较类
        // 这里first是nums[i]，second代表次数
        // 统计频率
        unordered_map<int,int>fre;
        for(auto &v:nums){
            fre[v]++;
        }
        for(auto &kv:fre){
            int value=kv.first;
            int freq=kv.second;
            if(minHeap.size()<k){
                minHeap.push(make_pair(value,freq));
            }else{
                if(minHeap.top().second<freq){
                    minHeap.pop();
                    minHeap.push(make_pair(value,freq));
                }
            }
        }
        vector<int> res;
        while(!minHeap.empty()){
            res.push_back(minHeap.top().first);
            minHeap.pop();
        }
        return res;
    }
};
