class Solution {
    /*
    有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。
示例 1：

输入：points = [[10,16],[2,8],[1,6],[7,12]]
输出：2
解释：气球可以用2支箭来爆破:
-在x = 6处射出箭，击破气球[2,8]和[1,6]。
-在x = 11处发射箭，击破气球[10,16]和[7,12]。
    题解：
    类似435最小会议数，相交的气球可以一支弓箭解决，故求的是不相交的
    气球右边坐标排序
    原因：https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/solutions/494602/tu-jie-tan-tao-wei-shi-yao-yao-an-qu-jian-de-you-d/?envType=study-plan-v2&envId=top-interview-150
    如果按左端排序，可能出现[0, 9], [0, 6], [7, 8]
    第一个区间和第二个重合，我让当前第一个区间继续寻求重合，它和第三个也重合。
    其实这时候第二个和第三个其实并不重合。
    如果按右端升序排序，就杜绝了「前面包后面」的情况
    当前区间x  -----R
      区间a ----------R
        区间b  L--------
    当前区间x遇到重合区间 a，a 的右端>=自己的右端（因为右端递增），x遇到下一个重合区间 b，b 的左端<=x的右端（因为有重合），所以a、b 一定有交集，参照物是当前区间x的右端。于是，放心地让当前区间一路找重合，直到遇到不重合，就有了一组能一箭穿的。
    这里示例1，按照右端排序后
   1 [1.......6] 
   2   [2..........8] 
   3             [7.........12] 
   4                 [10.........16]
    步骤：
    1. 拿当前区间1的右端作为标杆
    2. 只要 下一个区间的左端<=标杆下一个区间的左端<=标杆，则重合，继续寻求与下一个区间重合。此时区间2满足；
    3. 直到遇到不重合的区间，弓箭数 +1。这里区间3不满足；
    4. 拿新区间的右端作为标杆，也就是拿区间做标杆重复以上步骤。
    */
public:
    int findMinArrowShots(vector<vector<int>>& points) {
        if(points.size()<1){
            return 0;
        }
        // 气球右边坐标排序
        // 原因：https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/solutions/494602/tu-jie-tan-tao-wei-shi-yao-yao-an-qu-jian-de-you-d/?envType=study-plan-v2&envId=top-interview-150
        sort(points.begin(),points.end(),[](vector<int>&a,vector<int>&b){
            return a[1]<b[1];
        });
        int res=1;
        int curSegmentEnd=points[0][1];
        for(int i=1;i<points.size();i++){
            if(points[i][0]>curSegmentEnd) {
                // 遇到不重合的区间
                res++;
                curSegmentEnd=points[i][1];
            }
        }
        return res;
    }
};
