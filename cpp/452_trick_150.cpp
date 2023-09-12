class Solution {
    /*
    题目：
    有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
给你一个数组 points ，返回引爆所有气球所必须射出的最小弓箭数 。
    题解：
    类似435最小会议数，相交的气球可以一支弓箭解决，故求的是不相交的
    */
public:
    int findMinArrowShots(vector<vector<int>>& points) {
        if(points.size()<1){
            return 0;
        }
        // 按照气球的结束坐标排序
        sort(points.begin(),points.end(),[](vector<int>&a,vector<int>&b){
            return a[1]<b[1];
        });
        int res=1;
        int pre=points[0][1];
        for(int i=1;i<points.size();i++){
            if(points[i][0]>pre){
                // 不相交的条件：后一个气球的开始坐标大于（不能等于）前一个气球的结束坐标
                res++;
                pre=points[i][1];
            }
        }
        return res;
    }
};
