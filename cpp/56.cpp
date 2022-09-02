class Solution {
public:
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

*/
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>>res;
        if (intervals.size()<1){
            return res;
        }
        sort(intervals.begin(),intervals.end());
        for(int i=0;i<intervals.size();i++){
            int l=intervals[i][0];
            int r=intervals[i][1];
            if(res.empty()||res.back()[1]<l){
                // res 空 或者当前区间起始位置大于结果数组最后区间的结束位置（也即是不重叠）直接入
                res.push_back(intervals[i]);
            }else{
                // 重叠，结束位置取最大值
                res.back()[1]=max(res.back()[1],r);
            }
        }
        return res;
    }
};
