class Solution {
    /*
    给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须尽可能减少整个过程的操作步骤。
    相比较153，存在重复元素
    重复元素的处理是关键
    */
public:
    int findMin(vector<int>& nums) {
        int low=0;
        int high=nums.size()-1;
        while(low<high){
            int mid=low+(high-low)/2;
            if(nums[mid]>nums[high]){
                // 右侧
                low=mid+1;
            }else if(nums[mid]<nums[high]){
                // 左侧
                high=mid;
            }else{
                // 重复
                high--;
            }
        }
        return nums[low];
    }
};
