class Solution {
    /*
    题目：
    给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。
请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。
    思路：
    事实1:
    对于w: [1, 2, 3, 4]，可以将每个数值拆成1等份，则
    idx=0 只有1个等份；idx=1有2等份，idx=2有3个等份，idx=3有4个等份
    事实2:
    此外返回概率=w[i] / sum(w)
    故引入前缀和，sum[0] = W[0], sum[1] = W[0] + W[1]...
    求出一个随机数，next=rand() % sum.back(); 假设 sum.back() = 10, 那么这里产生的数字是 0-9。如果我们继续用以上的例子的话那么其每个数字所对应取到的index便为(结合概率要求)：
    next=0 ：小于sum[0],代表取到 index 0
    next=1，2: 小于sum[1],代表取到 index 1
    next=3，4，5: 小于sum[2],代表取到 index 2
    next=6，7, 8, 9: 小于sum[3],代表取到 index 3
    正好对应事实1里等份
    关键是怎么保证按概率返回，用以上的例子产生的前缀和表 [1, 3, 6, 10], 可以发现我们用得到的数字调用 upper_bound() 会刚好使其指向我们的 index 位置。
    next=0 的 upper_bound 会指向 index 0, 因为第一个比 0 大的数是 sum[0] = 1;
    next=1, 2 的 upper_bound 会指向 index 1, 因为第一个比 1 或者 2 大的数是 sum[1] = 3;
    next=3，4，5 的 upper_bound 会指向 index 2, 因为第一个比 {3, 4, 5} 大的数是 sum[2] = 6;
    next=6，7, 8, 9 的 upper_bound 会指向 index 3, 因为第一个比 {6，7, 8, 9} 大的数是 sum[3] = 10;
    */
    //w: [1, 2, 3, 4],前缀和表sum [1, 3, 6, 10],
public:
    Solution(vector<int>& w) {
        // cal prefix sum
        sum.push_back(w[0]);
        for(int i=1;i<w.size();i++){
            sum.push_back(sum.back()+w[i]);// push_back pop_back front back
        }
    }
    
    int pickIndex() {
        int next=random()%sum.back();//假设 w.back() = 10, 那么这里产生的数字是 0-9
        return upper_bound(sum.begin(),sum.end(),next)-sum.begin();
    }
private:
    vector<int>sum;// prefix sum
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(w);
 * int param_1 = obj->pickIndex();
 */
