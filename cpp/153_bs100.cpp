class Solution {
    /*
    给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [3,4,5,1,2]
输出：1
    题解：
    https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solutions/698479/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-5irwp/?envType=study-plan-v2&envId=top-100-liked
    我们考虑数组中的最后一个元素nums[high]：在最小值右侧的元素（不包括最后一个元素本身），
    它们的值一定都严格小于nums[high]；而在最小值左侧的元素，它们的值一定都严格大于nums[high]。
    因此，我们可以根据这一条性质，通过二分查找的方法找出最小值。
    */
public:
    int findMin(vector<int>& nums) {
        int low=0;
        int high=nums.size()-1;
        while(low<high){
            // 不能等于，否则[1]这种case会循环
            int mid=low+(high-low)/2;
            if(nums[mid]>nums[high]){
                // 说明 nums[mid]是最小值左侧的元素，因此我们可以忽略二分查找区间的左半部分
                low=mid+1;
            }else{
                // 说明 nums[mid]是最小值右侧的元素，因此我们可以忽略二分查找区间的右半部分
                high=mid;//[left,right)故mid可以取值
            }
        }
        return nums[low];
    }
};
