class Solution {
    /*
    题目等价从最左边抽取i个数字，从最右边抽取k-i个数字，能抽取的最大点数和多少
    也就是抽走的卡牌点数之和 = cardPoints 所有元素之和 - 剩余的中间部分元素之和
    1. 为快速求出中间部分元素之和，引入前缀和PreSum。preSum 方法能快速计算指定区间段 i ~ j 的元素之和
    preSum[0]=0;
    preSum[1]=preSum[0]+nums[0];
    ...
    preSum[i]=preSum[i-1]+nums[i-1];
    则sum(i,j)=preSum[j]-preSum[i-1];
    2.然后使用一个 0 ~ k 的遍历表示从左边拿走的元素数，然后根据窗口大小 windowSize = n - k ，利用 preSum 快速求窗口内元素之和
    3. 最大点数，也就是中间和最小
    */
public:
    const int int_max=0x7fffffff;
    int maxScore(vector<int>& cardPoints, int k) {
        int n=cardPoints.size();
        vector<int>preSum(n+1,0);
        preSum[0]=0;
        for(int i=1;i<=n;i++){
            preSum[i]=preSum[i-1]+cardPoints[i-1];
        }
        int res=int_max;//中间和的最小值
        int windowSize=n-k;
        for(int i=0;i<=k;i++){
            // 左边i个，右边k-i个，i取值0到k
            // 当左边取i个（也就是cardPoints[0:i-1]，也就是preSum[i]），右边k-i个，中间窗口是[i+1,i+windowSize]
            res=min(res,preSum[i+windowSize]-preSum[i]);
        }
        return preSum[n]-res;
    }
};
