class Solution {
    const int int_max=0x7FFFFFFF;
    /*
    给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
    请注意，顺序不同的序列被视作不同的组合。
    排列形式的背包问题，先总额再物品，这样这样才能让物品按照顺序放入背包
    参考：https://blog.csdn.net/qq_43530796/article/details/125410353
    完全背包问题
    dp[i][j]=max(dp[i-1][j],dp[i-1][j-k*w]+k*v)
    =max(dp[i-1][j],dp[i][j-w]+v)
    dp[0][j]=0;//没有元素，方案数为0
    dp[i][0]=1;//和为0，方案数为1
    */
public:
    int combinationSum4(vector<int>& nums, int target) {
        // 求解 顺序的完全背包 问题时，对 物品的迭代 应该放在最里层，对背包的迭代放在外层，只有这样才能让物品按一定顺序放入背包中
        // dp[i] 表示 总和为 i 的元素组合数，dp[i]+=dp[i−nums[j]]
        vector<int>dp(target+1,0);
        dp[0]=1;
        for(int j=1;j<=target;j++){
            // 先总额
            for(int i=1;i<=nums.size();i++){
                // 再物品
                if(j>=nums[i-1]){
                    //c++计算的中间结果会溢出，但因为最终结果是int
                    dp[j]=(0LL+dp[j]+dp[j-nums[i-1]])%int_max;
                }
            }
        }
        return dp[target];
    }
};
