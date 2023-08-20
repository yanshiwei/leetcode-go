class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        if(nums.size()<1){
            return 0;
        }
        int left=0;
        int right=nums.size()-1;//[left,right]
        while(left<=right){
            // 因为[left,right]，所以while这里等于，也就是left=right+1时候才结束
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                return mid;
            }else if(nums[mid]>target){
                // in left part
                right=mid-1;//// 因为[left,right] 所以取mid-1否则取mid时候重复
            }else{
                // in the eight
                left=mid+1;
            }
        }
        return left;
    }
};
