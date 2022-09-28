class Solution {
    /*
    一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字
    */
public:
    int missingNumber(vector<int>& nums) {
        if(nums.size()<1){
            return -1;
        }
        int left=0, right=nums.size()-1;
        while(left<=right){
            int midIdx=(left+right)>>1;
            if(nums[midIdx]==midIdx){
                // 相等则说明缺失的数字在右侧
                left=midIdx+1;
            }else{
                // 不相等在左侧
                right=midIdx-1;
            }
        }
        return left;
    }
};
