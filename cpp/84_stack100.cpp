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
    一个坐标所能构成的最大面积时，只需要看左右侧第一个比它小的值即可（起始点终点外侧默认为0），很容易想到单调栈
遍历的时候，需要记录的是下标，如果当前的高度比它之前的高度严格小于的时候，就可以直接确定之前的那个高的柱形的最大矩形的面积，为了确定这个最大矩形的左边界，我们还要找到第一个严格小于它的高度的矩形，向左回退的时候，其实就可以当中间这些柱形不存在一样。
    */
public:
    int largestRectangleArea(vector<int>& heights){
        /*
        当前数字向两边扩展，遇到比本身大的就接着扩展直到遇到比自己小的。然后用自己的高度乘以拓展的宽度，每次都将求的的面积跟最大面积res比较。
        hs[i] 作为矩形高度时,l[i]代表位置 i左边最近一个比其小的位置，r[i]代表位置 i右边最近一个比其小的位置，那么 r[i]−l[i]−1则是以 hs[i]作为矩形高度时所能取得的最大宽度
        “比本身大的就接着扩展直到遇到比自己小的”这就是单调递减栈特性（从栈低到栈顶数值从小到大）
        本题目关键是找到每个柱子左右两边第一个小于本身的柱子，对应与单调递减栈。
        其实就是栈顶元素，和栈顶元素的前一个元素（左边第一个小于的元素）以及将要入栈的元素（右边第一个小于的元素）组成了该栈顶元素对应的面积最大值。
        举例：heights = [2,1,5,6,2,3]
        i=0时，左侧无最小（l=-1），右侧最小值位置是1故宽度为1，高arr[0]=2
        i=1时，左侧无最小（l=-1），右侧无最小(r=size)故宽度为6，高arr[1]=1
        i=2时，左侧最小值位置是1，右侧最小值位置是4故宽度为2，高arr[2]=5
        i=3时，左侧最小值位置是2，右侧最小值位置是4故宽度为1，高arr[3]=6
        i=4时，左侧最小值位置是1，右侧无最小(r=size)故宽度为4，高arr[4]=2
        i=4时，左侧最小值位置是4，右侧无最小(r=size)故宽度为1，高arr[5]=3
        */
        int res=0;
        heights.push_back(0);//为了确保全部计算，这样栈内元素可以全部计算包括最后一个元素
        stack<int>min_st;//从堆底到堆顶从小到大
        for(int i=0;i<heights.size();i++){
            // 单调递减栈,如果栈为空或入栈元素值大于栈顶元素值则入栈
            if(min_st.empty()||heights[min_st.top()]<heights[i]){
                min_st.push(i);
            }else{
                // 如果入栈则会破坏栈的单调性，则需要把比入栈元素大的元素全部出栈
                while(!min_st.empty()&&heights[min_st.top()]>=heights[i]){
                    int idx=min_st.top();
                    min_st.pop();
                    int left=min_st.empty()?-1:min_st.top();//左边第一个小于的元素
                    int right=i;//右边第一个小于的元素
                    int w=right-left-1;
                    int h=heights[idx];
                    res=max(res,w*h);
                }
                min_st.push(i);
            }
        }
        return res;
    }
    /*
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
    */
};
