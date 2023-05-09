class Solution {
    /*
    题目：
    这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
注意，时间上出现重叠的 2 份工作不能同时进行。
如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。
    题解：
    1.工作按照工作结束时间从小到大排序，前 i个工作的最大报酬
        1.1 不选第 i 个工作，那么最大报酬等于前 i−1个工作的最大报酬（转换成了一个规模更小的子问题）；
        1.2 选第 i个工作，由于工作时间不能重叠，设 j是最大的满足 endTime[j]≤startTime[i]的j，那么最大报酬等于前 j个工作的最大报酬加上 profit[i]（同样转换成了一个规模更小的子问题）
        1.3 这两种决策取最大值
    2.定义 dp[i]表示按照结束时间排序后的前 i 个工作的最大报酬，其中i对应profit/startTime里的i-1下标，对应分类讨论：
        2.1 不选第 i 个工作，dp[i]=dp[i-1];
        2.2 选第 i个工作，dp[i]=dp[j]+profit[i-1]
        2.3 dp[i]=max(dp[i-1],dp[j]+profit[i-1])
        2.4 初始化：dp[0]=0,0个任务收获为0
        2.5 最终结果是dp[n]
    3 求最大的满足 endTime[j]≤startTime[i]的j，可以用upper_bound
    对一个有序序列的一个左闭右开的有序区间进行二分查找，区间左端点和右端点分别为第1、2个参数，第三个参数是查找的值，第四个参数可填一个cmp函数，也可不填。
返回容器中第一个大于目标值的位置。若容器中的元素都比目标值小则返回最后一个元素的下一个位置。
    */
public:
    int jobScheduling(vector<int>& startTime, vector<int>& endTime, vector<int>& profit) {
        int n=startTime.size();
        vector<vector<int>>jobs(n);
        for(int i=0;i<n;i++){
            jobs[i]={startTime[i],endTime[i],profit[i]};
        }
        // 按结束时间排序
        sort(jobs.begin(),jobs.end(),[](vector<int>&a,vector<int>&b)->bool{
            return a[1]<b[1];
        });
        vector<int>dp(n+1);//dp[i]前i个任务获得的最大收获，第i个任务对应profit/startTime里的i-1下标
        dp[0]=0;//init,0个任务收获为0
        for(int i=1;i<=n;i++){
            int j = upper_bound(jobs.begin(), jobs.begin() + i - 1, jobs[i - 1][0], [](int st, const vector<int> &job) -> bool {
                return st < job[1];
            }) - jobs.begin();
            // 为什么是 j 不是 j-1：上面算的是 > 开始时间也就是profit/startTime里的下标，j-1 后得到 <= 开始时间，但由于dp里下标关系还要 +1，抵消了
            dp[i]=max(dp[i-1],dp[j]+jobs[i-1][2]);
        }
        return dp[n];
    }
};
