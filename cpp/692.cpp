class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        unordered_map<string,int>cnt;//insert erase size empty find count
        for(auto&word:words){
            cnt[word]++;
        }
        // TOPK 小顶堆
        auto cmp=[](const pair<string,int>&a,const pair<string,int>&b){
            // small heap so greate
            if(a.second==b.second){
               //如果不同的单词有相同出现频率， 按字典顺序 排序
               return a.first<b.first; // 标准2 更大一点 字典顺序更靠前
            }
            return a.second>b.second;// 标准1 更大一点 次数更大
        };
        // push pop top size empty
        priority_queue<pair<string,int>,vector<pair<string,int>>,decltype(cmp)>q(cmp);
        for(auto&it:cnt){
            q.push(it);
            if(q.size()>k){
                q.pop();
            }
        }
        vector<string>res(k);
        for(int i=k-1;i>=0;i--){
            res[i]=q.top().first;
            q.pop();
        }
        return std::move(res);
    }
};
