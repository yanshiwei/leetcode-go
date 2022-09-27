class Solution {
public:
    // class重载()操作符，定义了一个函数对象
    class cmp{
        public:
          bool operator()(pair<int, int>& m, pair<int, int>& n){
              return m.second>n.second;
          }
    };
    vector<int> topKFrequent(vector<int>& nums, int k) {
        if(nums.size()<1||k<1){
            return {};
        }
        unordered_map<int,int>occ;// insert erase find count empty size
        for(auto &v:nums){
            occ[v]++;
        }
        // pair 的第一个元素代表数组的值，第二个元素代表了该值出现的次数
        priority_queue<pair<int,int>,vector<pair<int,int>>,cmp>q;// like queue: push pop top size empty swap
        for(auto&[num,cnt]:occ){
            if(q.size()==k){
                if(q.top().second<cnt){
                    q.pop();
                    q.emplace(num,cnt);
                }
            }else{
                q.emplace(num,cnt);
            }
        }
        vector<int>res;
        while(q.empty()==false){
            res.emplace_back(q.top().first);
            q.pop();
        }
        return res;
    }
};
