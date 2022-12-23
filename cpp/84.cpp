class Solution {
    /*
    题目：
    给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。求在该柱状图中，能够勾勒出来的矩形的最大面积。
    题解：
    https://leetcode.cn/problems/largest-rectangle-in-histogram/solutions/142012/bao-li-jie-fa-zhan-by-liweiwei1419/
    heights = [2,1,5,6,2,3]
    i=0 area=2*1
    i=1 area=1*2,h[1]<h[0]只取较小值
    i=2 area=5*1
    i=3 area=5*2
    i=4 area=2*3
    i=5 area=2*4
    故秋一个坐标所能构成的最大面积时，只需要看左右侧第一个比它小的值即可（起始点终点外侧默认为0），很容易想到单调栈
    */
public:
    int largestRectangleArea(vector<int>& heights) {
        if(heights.size()<1){
            return 0;
        }
        stack<int>st;//store idx 单调递增
        int res=0;
        int i=0;
        while(i<heights.size()){
            if(st.empty()){
                st.push(i);
                i++;
            }else{
                int top=st.top();
                if(heights[i]>=heights[top]){
                    st.push(i);
                    i++;
                }else{
                    int h=heights[st.top()];
                    st.pop();
                    //左边第一个小于当前柱子的下标，因为栈内从小到大排序
                    int left=st.empty()?-1:st.top();
                    //右边第一个小于当前柱子的下标
                    int right=i;
                    res=max(res,h*(right-left-1));
                    
                }
            }
        }
        while(st.empty()==false){
            int h=heights[st.top()];
            st.pop();
            //左边第一个小于当前柱子的下标，因为栈内从小到大排序
            int left=st.empty()?-1:st.top();
            //右边没有小于当前高度的柱子，所以赋值为数组的长度便于计算，栈内都是小值，其面积可以从left延伸到结尾
            int right=heights.size();
            res=max(res,h*(right-left-1));
        }
        return res;
    }
};
