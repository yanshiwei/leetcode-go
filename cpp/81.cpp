class Solution {
    /*
    已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
    */
public:
    bool search(vector<int>& nums, int target) {
        if(nums.size()<1){
            return false;
        }
        int left=0;
        int right=nums.size()-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                return true;
            }
            if(nums[mid]==nums[left]){
                left++;
                continue;
            }
            if(nums[mid]>nums[left]){
                // 左边有序
                if(target>=nums[left]&&target<nums[mid]){
                    right=mid-1;
                }else{
                    left=mid+1;
                }
            }else{
                if(nums[mid]==nums[right]){
                    right--;
                    continue;
                }
                // 右边有序
                if(nums[mid]<target&&target<=nums[right]){
                    left=mid+1;
                }else{
                    right=mid-1;
                }
            }
        }
        return false;
    }
};
