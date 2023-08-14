class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>> res;
        if(intervals.size()<1){
            return res;
        }
        // 按照区间的左端点排序
        sort(intervals.begin(),intervals.end());
        for(int i=0;i<intervals.size();i++){
            int l=intervals[i][0];
            int r=intervals[i][1];
            if(res.empty()){
                // empty
                res.push_back(intervals[i]);
            }else if(res.back()[1]<l){
                // 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组 merged 的末尾
                res.push_back(intervals[i]);
            }else{
                // overlap
                // 重合，我们需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，将其置为二者的较大值
                res.back()[1]=max(res.back()[1],r);
            }
        }
        return res;
    }
};
