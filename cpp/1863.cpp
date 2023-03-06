class Solution {
    /*
    一个数组的 异或总和 定义为数组中所有元素按位 XOR 的结果；如果数组为 空 ，则异或总和为 0 。
例如，数组 [2,5,6] 的 异或总和 为 2 XOR 5 XOR 6 = 1 。
给你一个数组 nums ，请你求出 nums 中每个 子集 的 异或总和 ，计算并返回这些值相加之 和 。
    */
public:
    int subsetXORSum(vector<int>& nums) {
        int res=0;
        for(int i=0;i<(1<<nums.size());i++){
            // 二进制模拟所有子集
            int tmpSum=0;
            for(int j=0;j<nums.size();j++){
                if(i&(1<<j)){
                    tmpSum^=nums[j];
                }
            }
            res+=tmpSum;
        }
        return res;
    }
};
