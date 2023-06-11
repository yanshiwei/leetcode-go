class Solution {
    /*
    给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。
如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。
    */
public:
    bool checkSubarraySum(vector<int>& nums, int k) {
        int n=nums.size();
        vector<int>preSum(n+1,0);
        // 求前缀和
        preSum[0]=0;
        for(int i=1;i<=n;i++){
            preSum[i]=preSum[i-1]+nums[i-1];
        }
        // 求[left,right]的区间和,left/right是nums下标，转为preSum下标为left+1/right+1
        // 前right+1个项目和-前left个项目和
        // sum(left,right)=preSum[right+1]-preSum[left]
        // 固定right的时候移动left遍历每种情况
        for(int right=1;right<n;right++){
            // 为了保证子数组长度至少2，故终点从1开始
            // 前缀和比k小跳过，当有两个连续的0时除外
            if(preSum[right+1]<k&&preSum[right+1]!=preSum[right-1]){
                // preSum[right]==preSum[right-1]说明nums[right-1],nums[right]两个连续的0，两个连续0和为0也是k倍数
                continue;
            }
            // 遍历每种左端点，子数组大小 至少为 2 故距离至少为2
            // 也就是right-left+1>=2====>left<=right-1
            for(int left=0;left<=right-1;left++){
                if((preSum[right+1]-preSum[left])%k==0){
                    return true;
                }
            }
        }
        return false;
    }
};
