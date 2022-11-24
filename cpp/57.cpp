class Solution {
    /*
    题目：
    给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
    */
public:
    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        vector<vector<int>> res;
        if(intervals.size()<1){
            res.push_back(newInterval);
            return res;
        }
        //找到插入的位置，也就是找到第一个满足intervals[i].end>=newInterval.start(看newInterval左侧)
        int idx=0;
        while(idx<intervals.size()&&intervals[idx][1]<newInterval[0]){
            res.push_back(intervals[idx]);
            idx++;
        }
        // 合并重叠区域，直到newInterval与intervals[i]不重叠，
        // 也就是intervals[i].start>newInterval.end(看newInterval右侧)
        while(idx<intervals.size()&&intervals[idx][0]<=newInterval[1]){
            newInterval={min(intervals[idx][0],newInterval[0]),max(intervals[idx][1],newInterval[1])};
            idx++;
        }
        // 插入
        res.push_back(newInterval);
        // 插入剩下的
        while(idx<intervals.size()){
            res.push_back(intervals[idx]);
            idx++;           
        }
        return res;
    }
};
