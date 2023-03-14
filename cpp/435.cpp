class Solution {
    /*
    给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。
    贪心：预定会议问题。
    */
public:
    int eraseOverlapIntervals(vector<vector<int>>& intervals) {
        // 求最少的移除区间数，也就是尽量多的保留不重叠区间
        // 在选择保留区间时，区间的结尾越小，预留给其他区间的空间就越大就能保留更多区间
        // 采取的贪心策略为：优先保留结尾小且不相交的区间。
        // 具体实现方法为，先把区间按照结尾的大小进行增序排序，
        // 每次选择结尾最小且和前一个选择的区间不重叠的区间
        if(intervals.size()<2){
            return 0;
        }
        sort(intervals.begin(),intervals.end(),[](vector<int>&a,vector<int>&b){
            return a[1]<b[1];//区间按照结尾的大小进行增序排序
        });
        int prev=intervals[0][1];
        int cnt=0;
        for(int i=1;i<intervals.size();i++){
            if(prev>intervals[i][0]){
                //如果下一个区间的开头小于上一个区间的结尾，证明两区间重叠，需要移除该区间
                cnt++;
            }else{
                prev=intervals[i][1];
            }
        }
        return cnt;
    }
};
