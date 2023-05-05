class Solution {
    /*
    给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
    low与mid比较？
    nums[low]<nums[mid]：最小值在左侧；
    nums[low]>nums[mid]：最小值在左侧；
    比较不出什么；
    high与mid比较？
    nums[high]<nums[mid]：最小值在右侧；
    nums[high]>nums[mid]：最小值在左侧；   
    故使用high与mid比较
    */
public:
    int findMin(vector<int>& nums) {
        int low=0;
        int high=nums.size()-1;
        while(low<high){
            int mid=low+(high-low)/2;
            if(nums[mid]>nums[high]){
                // 右侧，由于是low<high结束，为了避免死循环
                low=mid+1;
            }else{
                // 左侧
                high=mid;
            }
        }
        return nums[low];
    }
};
