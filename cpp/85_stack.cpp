class Solution {
    /*
    题目：
    给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
    题解：
    相比较84题，这是二维，首先是将二维转为一维。
    具体来说：
    从第一行到第n行形成的柱状图可以利用84题求解，然后循环每一行，计算以这一行为底的柱状图最大面积，然后更新最大矩形面积
    举例清参考：
    https://leetcode.cn/problems/maximal-rectangle/solutions/1774721/by-burling-lucky-gqg0/
    */
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        int m=matrix.size();
        if(m<1){
            return 0;
        }
        int n=matrix[0].size();
        if(n<1){
            return 0;
        }
        vector<int>heights(n,0);//每一行的个数都是0
        int maxAres=0;
        for(int i=0;i<m;i++){
            //针对每一行，构造heights
            // 不能 heights.clear()，需要递增上一行结果
            for(int j=0;j<n;j++){
                if(matrix[i][j]=='0'){
                    heights[j]=0;//中间有0就重置
                }else{
                    heights[j]+=1;
                }
            }
            maxAres=max(maxAres,largestRectangleArea(heights));
        }
        return maxAres;
    }
private:
    int largestRectangleArea(vector<int>&heights){
        // 84题解法，求高度围成的最大面积,单调递减栈
        stack<int>min_st;
        heights.push_back(0);//为了确保全部计算，包括最后一个元素
        int res=0;
        for(int i=0;i<heights.size();i++){
            // 单调递减栈,如果栈为空或入栈元素值大于栈顶元素值则入栈
            if(min_st.empty()||(heights[min_st.top()]<heights[i])){
                min_st.push(i);
            }else{
                // 如果入栈则会破坏栈的单调性，则需要把比入栈元素大的元素全部出栈
                while(!min_st.empty()&&heights[min_st.top()]>heights[i]){
                    int mid=min_st.top();
                    min_st.pop();
                    int left=min_st.empty()?-1:min_st.top();//左边第一个小于的元素
                    int right=i;//右边第一个小于的元素
                    int w=right-left-1;
                    int h=heights[mid];
                    res=max(res,w*h);
                }
                min_st.push(i);
            }
        }
        return res;
    }
};
