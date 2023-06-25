class Solution {
    /*
    题目：
    给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
    题解：
    类似347，统计频率最高TOPK，求TOPK用最小堆。不同的是本题要求频率一致情况下按照字典顺序
    构造大小k的最小堆，当堆顶元素小于当前值，替换；当堆顶元素大于当前值，跳过
    */
public:
    struct greater{
        //词频由高到低排序：最小堆；字母顺序排序：最大堆
        bool operator()(pair<string,int>&m,pair<string,int>&n){
            if(m.second!=n.second){
                return m.second>n.second;
            }
            return m.first<n.first;
        }
    };
    vector<string> topKFrequent(vector<string>& words, int k) {
        if(k<1){
            return {};
        }
        unordered_map<string,int>fre;
        for(auto&w:words){
            fre[w]++;
        }
        priority_queue<pair<string,int>,vector<pair<string,int>>,greater>mihHeap;
        for(auto kv:fre){
            string w=kv.first;
            int cnt=kv.second;
            mihHeap.push(make_pair(w,cnt));
            // 维护大小k的最小堆
            if(mihHeap.size()>k){
                mihHeap.pop();//推出较小值
            }
        }
        vector<string> res;
        while(!mihHeap.empty()){
            res.push_back(mihHeap.top().first);
            mihHeap.pop();
        }
        // 返回的答案应该按单词出现频率由高到低排序，而res是按照词频从低到高所以reverse
        reverse(res.begin(),res.end());
        return res;
    }
};
