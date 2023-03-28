class Solution {
    /*
    题目：
    给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。0 <= nums[i] <= 1000
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
    题解：
    01背包问题，https://leetcode.cn/problems/target-sum/solutions/1410230/by-flix-rkb5/
    1 01背包问题是选或者不选，但本题是必须选，是选+还是选-。先将本问题转换为01背包问题。
        设数组和为total，（0 <= nums[i] <= 1000），要想得到target，肯定是正负组合如例子里nums = [1,1,1,1,1], target = 3
        则一种组合是-1 + 1 + 1 + 1 + 1 = 3，此时正数集合下标pos={1，2，3，4},负数集合下标neg={0}
        此时pos-nes=target
        而。pos+nes=totaol
        则pos=(total+target)/2;
         neg=(total-target)/2;
         此时不难发现，本题实质上是一道「0-1 背包问题」：给定一个只包含正整数的非空数组 nums，判断是否可以从数组中选出一些数字（每个元素最多选一次），使得选出的这些数字的和刚好等于 pos或者 neg
    2 对于本题而言，nums[i] 则对应于常规背包问题中第 i件物品的重量。我们要做的是从数组nums 中选出若干个数字（每个元素最多选一次）使得其和刚好等于 pos或者 neg，并计算有多少种不同的选择方式。
    2.1 状态定义：dp[i][j] 表示：从前 i个数字中选出若干个，使得被选出的数字其和为 j的方案数目。
    2.2 状态转移：dp[i][j]=dp[i−1][j] + dp[i−1][j−nums[i]]（不考虑这个物品或者考虑）
    2.3 初始化：dp[0][0]=1：表示从前 0个数字中选出若干个数字使得其和为 0的方案数为 1，即「空集合」不选任何数字即可得到 0。
对于其他 dp[0][j],  j≥1，则有 dp[0][j]=0：「空集合」无法选出任何数字使得其和为 j≥1。
    3. 最后根据01背包问题降纬
    */
public:
    int findTargetSumWays(vector<int>& nums, int target) {
        int total=accumulate(nums.begin(),nums.end(),0);
        if (abs(target)>total){
            // 否则neg=(total-target)/2负数 不可能得到
            return 0;
        }
        if((total+target)%2!=0){
            //(total+target)/2不是整数，不可能得到
            return 0;
        }
        int pos=(total+target)/2;
        int neg=(total-target)/2;
        int cap=min(pos,neg);// 取pos和neg中的较小者，以使得dp空间最小
        vector<int>dp(cap+1,0);//代表第i个元素对于nums[i-1]
        // dp[j]表示容量为j的背包能放下东西的最大价值这里代表和为j时方案数目
        dp[0]=1;
        for(int i=0;i<nums.size();i++){
            // 01背包逆序
            for(int j=cap;j>=nums[i];j--){
                dp[j]+=dp[j-nums[i]];
            }
        }
        return dp[cap];
    }
};
