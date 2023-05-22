class Solution {
    /*
    题目：
    你需要制定一份 d 天的工作计划表。工作之间存在依赖，要想执行第 i 项工作，你必须完成全部 j 项工作（ 0 <= j < i）。
你每天 至少 需要完成一项任务。工作计划的总难度是这 d 天每一天的难度之和，而一天的工作难度是当天应该完成工作的最大难度。
给你一个整数数组 jobDifficulty 和一个整数 d，分别代表工作难度和需要计划的天数。第 i 项工作的难度是 jobDifficulty[i]。
返回整个工作计划的 最小难度 。如果无法制定工作计划，则返回 -1 。
    题解：
    dp[i][j]前i天完成前j项工作的最小难度(i从0开始)，前i天完成前j项任务， 它有多种选择
    1. 前i-1天完成k-1项任务，因为每天至少完成一项，所以k-1>=i-1(前i-1前每天一个),然后第i天完成k,k+1,...,j项任务，而一天的难度是所有工作的最大值故第i天难度f(i,j)=max{jobDifficulty[k],jobDifficulty[k+1],...,jobDifficulty[j]}，最终所需难度
    dp[i-1][k-1]+max{jobDifficulty[k],jobDifficulty[k+1],...,jobDifficulty[i]}。 
    其中i=k<=jobDifficulty.size()
    为了完成的难度最低，我们从中选最小的即可：
    dp[i][j]=min{dp[i-1][k-1]+max{jobDifficulty[k],jobDifficulty[k+1],...,jobDifficulty[i]}}
    其中i=k<=jobDifficulty.size()
    */
    const int int_max=0x7fffffff;
public:
    int minDifficulty(vector<int>& jobDifficulty, int d) {
        if(jobDifficulty.size()<d){
            return -1;
        }
        vector<vector<int>>dp(d,vector<int>(jobDifficulty.size(),int_max));
        // init 第0天完成前j项工作的最小难度,一天的难度是所有工作的最大值
        int preDif=0;
        for(int j=0;j<jobDifficulty.size();j++){
            preDif=max(preDif,jobDifficulty[j]);
            dp[0][j]=preDif;
        }
        for(int i=1;i<d;i++){
            // 前第i天 至少有i个工作
            for(int j=i;j<jobDifficulty.size();j++){
                // 前j项任务
                // 前i-1天完成k-1项任务，第i天完成k,k+1,...,j项任务
                // 一天的难度是所有工作的最大值故第i天难度f(i,j)=max{jobDifficulty[k],jobDifficulty[k+1],...,jobDifficulty[j]}
                // i<=k<=j;
                preDif=jobDifficulty[j];
                for(int k=j;k>=i;k--){
                    preDif=max(preDif,jobDifficulty[k]);
                    dp[i][j]=min(dp[i][j],dp[i-1][k-1]+preDif);
                }
            }
        }
        return dp[d-1][jobDifficulty.size()-1];
    }
};
