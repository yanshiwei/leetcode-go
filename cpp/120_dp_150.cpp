class Solution {
    /*
    给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
  输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。  
题解：
动态规划
dp[i][j]代表第i行第j个最小路径和,i>=0,j>=0,j<=i(三角形特性保证)
1 转化公式
dp[i][j]=min(dp[i-1][j-1],dp[i-1][j])+t[i][j];
2 init
dp[i][0]=dp[i-1][0]+t[i][0];//对于第i行最左侧即j=0时，按照题意只能是上一行最左侧即转化公式为
dp[i][i]=dp[i-1][i-1]+t[i][i];//对于第i行最右侧即j=i时，按照题意只能是上一行前一位
dp[0][0]=t[0][0];
    */
public:
    int int_max=0x7fffffff;
    int getMin(vector<int>&nums){
        int res=nums[0];
        for(int i=1;i<nums.size();i++){
            res=min(res,nums[i]);
        }
        return res;
    }
    int minimumTotal(vector<vector<int>>& triangle) {
        int m=triangle.size();
        if(m<1){
            return 0;
        }
        vector<vector<int>>dp(m,vector<int>(m,0));
        // init
        dp[0][0]=triangle[0][0];
        for(int i=1;i<m;i++){
            //第i行长度为i+1,i>=0
            //最左侧处理
            dp[i][0]=dp[i-1][0]+triangle[i][0];
            // 非两端的中间部分处理
            for(int j=1;j<i;j++){
                dp[i][j]=min(dp[i-1][j],dp[i-1][j-1])+triangle[i][j];
            }
            //最右侧处理
            dp[i][i]=dp[i-1][i-1]+triangle[i][i];
        }
        return getMin(dp[m-1]);
    }
};
