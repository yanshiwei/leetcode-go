class Solution {
public:
    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        vector<vector<int>> res;
        if(intervals.size()<1){
            res.push_back(newInterval);
            return res;
        }
        // 找到插入的位置，看newInterval.start，也就是
        // 找到第一个满足intervals[i].end>=newInterval.start(看newInterval左侧)
        int idx=0;
        while(idx<intervals.size()&&intervals[idx][1]<newInterval[0]){
            res.push_back(intervals[idx]);//intervals[idx] 与 newInterval不交叠
            idx++;
        }
        // 合并重叠区域，直到newInterval与intervals[i]不重叠，看newInterval.end
        // 也就是intervals[i].start>newInterval.end(看newInterval右侧)
        while(idx<intervals.size()&&intervals[idx][0]<=newInterval[1]){
            newInterval={min(intervals[idx][0],newInterval[0]),max(intervals[idx][1],newInterval[1])};
            idx++;
        }
        // 插入
        res.push_back(newInterval);
        // 处理剩下的
        while(idx<intervals.size()){
            res.push_back(intervals[idx]);
            idx++;
        }
        return res;
    }
};
