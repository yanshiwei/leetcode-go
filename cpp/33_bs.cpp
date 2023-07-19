class Solution {
public:
    int search(vector<int>& nums, int target) {
        if(nums.size()<1){
            return -1;
        }
        if(nums.size()==1){
            return nums[0]==target?0:-1;
        }
        int left=0;
        int right=nums.size()-1;//[left,right]
        while(left<=right){
            int midIdx=left+(right-left)/2;
            if(nums[midIdx]==target){
                return midIdx;
            }
            if(nums[left]<=nums[midIdx]){
                //[left,midIdx-1] is sorted
                if(nums[left]<=target&&target<=nums[midIdx]){
                    // in [left,midIdx-1]
                    right=midIdx-1;
                }else{
                    // in [midIdx+1,right]
                    left=midIdx+1;
                }
            }else{
                //[midIdx+1,right] is sorted
                if(nums[midIdx+1]<=target&target<=nums[right]){
                    // in [midIdx+1,right] 
                    left=midIdx+1;
                }else{
                    // in [left,midIdx-1]
                    right=midIdx-1;
                }
            }
        }
        return -1;
    }
};
