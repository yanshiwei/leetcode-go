class Solution {
    /*
    整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1
    */
public:
    int search(vector<int>& nums, int target) {
        if(nums.size()<1){
            return -1;
        }
        if(nums.size()==1){
            return nums[0]==target?0:-1;
        }
        int left=0,right=nums.size()-1;
        while(left<=right){
            int midIdx=left+(right-left)/2;
            if(nums[midIdx]==target){
                return midIdx;
            }
            if(nums[left]<=nums[midIdx]){
                //[left,midIdx] is sorted
                if(nums[left]<=target&&target<=nums[midIdx]){
                    // in left
                    right=midIdx-1;
                }else{
                    // in right
                    left=midIdx+1;
                }
            }else{
                //[mid+1,right] is sorted
                if(target>nums[midIdx]&&target<=nums[right]){
                    // in right
                    left=midIdx+1;
                }else{
                    // in left
                    right=midIdx-1;
                }
            }
        }
        return -1;
    }
};
