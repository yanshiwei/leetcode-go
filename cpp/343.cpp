class Solution {
    /*
    题目：
    给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
返回 你可以获得的最大乘积 。
    */
public:
    int integerBreak(int n) {
        // dp[i]代表i拆成至少2个数后，这些数乘最大值
        // i=0,1 时不能拆分，所有dp[i]=0
        // i>=2时，可以拆分，设第一个拆分出的是j(1<=j<i)，则存在两种方案：
        // 1. i拆成j和i-j，且i-j不再拆分，此时乘j*(i-j);
        // 2. i拆成j和i-j，且i-j再拆分，此时乘j*dp[i-j];
        // 故j固定时，dp[i]=max(j*(i-j),j*dp[i-j])。由于 j的取值范围是 1到i−1 ，需要遍历所有的 j 得到 dp[i]的最终值
        vector<int> dp(n+1);
        dp[0]=0;
        dp[1]=0;
        dp[2]=1;
        for(int i=3;i<=n;i++){
            for(int j=1;j<i;j++){
                dp[i]=max(dp[i],max(j*(i-j),j*dp[i-j]));
            }
        }
        return dp[n];
    }
};
