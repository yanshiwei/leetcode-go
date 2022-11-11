class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        if(nums.size()<1){
            return -1;
        }
        int left=0;
        int right=nums.size()-1;
        while(left<=right){
            int mid=(left+(right-left)/2);
            if(nums[mid]==target){
                return mid;
            }
            if(nums[mid]>target){
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        if(right<0){
            return 0;
        }
        if(target>nums[right]){
            return right+1;
        }
        return -1;
    }
};
