class Solution {
    /*
    n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
    */
public:
    int candy(vector<int>& ratings) {
        // 贪心
        int n=ratings.size();
        if(n<2){
            return n;
        }
        // 把所有孩子的糖果数初始化为 1
        vector<int>candy(n,1);
        // 从左往右：如果右边孩子的评分比左边的高，则右边孩子的糖果数更新为左边孩子的糖果数加 1
        for(int i=0;i<n-1;i++){
            if(ratings[i+1]>ratings[i]){
                candy[i+1]=candy[i]+1;
            }
        }
        // 从右往左：如果左边孩子的评分比右边的高，且左边孩子当前的糖果数不大于右边孩子的糖果数，
        // 则左边孩子的糖果数更新为右边孩子的糖果数加 1
        // 要求左边孩子当前的糖果数不大于右边孩子的糖果数 是避免值被覆盖 like [1,3,4,5,2]
        // 经过从左到右后 candy=[1,2,3,4,1]
        // 如果不加这个条件，经过从右到左后， candy=[1,2,3,2,1]
        for(int i=n-1;i>0;i--){
            if(ratings[i-1]>ratings[i]&&candy[i-1]<=candy[i]){
                candy[i-1]=candy[i]+1;
            }
        }
                for(int i=0;i<n;i++){
            cout<<candy[i]<<",";
        }
        return accumulate(candy.begin(),candy.end(),0);//累加
    }
};
