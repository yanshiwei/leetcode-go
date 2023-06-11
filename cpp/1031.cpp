class Solution {
    /*
    https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/solutions/459960/qian-zhui-he-bian-li-onshi-jian-fu-za-du-ji-ke-qiu/
    对于有两个变量的题目，通常可以枚举其中一个变量，把它视作常量，从而转化成只有一个变量的问题。
    例如一个数组，求里面最大的两个数的和，那么我们可以遍历一次数组，每次固定第2个数A[i]的同时，去这个数的左边求第1个数的最大值max(A[0,1,...,i-1])；然后max(A[i]+max(A[0,1,...,i-1]))就是最终的结果
    回到这一题，不是求两个数的最大和，而是两个数组的最大和。由于两个数组L和M的长度是固定的，所以我们可以沿用前面的思路：
    1. 每次固定M的位置，L有多个位置，移动L求sumL的最大值，然后最大和就是max(sumM+max(sumL))
    2. 问题关键如何快速求出sumL，由此引入了前缀和
    */
public:
    int maxSumTwoNoOverlap(vector<int>& nums, int firstLen, int secondLen) {
        int n=nums.size();
        int maxL=0;
        int maxM=0;
        int maxLM=0;
        int maxML=0;
        vector<int>preSum(n+1,0);// preSum[i]代表前i个，i=0代表空
        vector<int>sumL(n+1,0);//sumL[i]为i为起始点，长度firstLen的区间和,i=0代表nums[0]开始
        vector<int>sumM(n+1,0);//sumM[i]为i为起始点，长度secondLen的区间和,i=0代表nums[0]开始
        //  1、先求数组A的前缀和preSum 2、利用preSum可以求出每一个以i为终点的sumL和sumM
        for(int i=1;i<=n;i++){
            preSum[i]=preSum[i-1]+nums[i-1];
            // 求出i为终点，长度为firstLen的区间和
            if(i>=firstLen){
                sumL[i-firstLen]=preSum[i]-preSum[i-firstLen];
            }
            // 求出i为终点，长度为secondLen的区间和
            if(i>=secondLen){
                sumM[i-secondLen]=preSum[i]-preSum[i-secondLen];
            }
        }
        // 每次固定M的位置，L有多个位置，移动L求sumL的最大值，然后最大和就是max(sumM+max(sumL))
        // 先假设L在左边，M在右边，L的终点位置可以从firstLen到n-secondLen，每次固定sumM[i]，找出最大的L
        for(int i=firstLen;i<=n-secondLen;i++){
            // i-1为L的终点，固定M位置在i，故i又是M的起点
            maxL=max(maxL,sumL[i-firstLen]);
            maxLM=max(maxLM,maxL+sumM[i]);
        }
        // 再假设M在左边，L在右边
        maxL = maxM = 0;
        for(int i=secondLen;i<=n-firstLen;i++){
            maxM=max(maxM,sumM[i-secondLen]);
            maxML=max(maxML,maxM+sumL[i]);
        }
        return max(maxLM,maxML);
    }
};
