class Solution {
    /*
    给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
    */
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        if(nums.size()<1){
            return {-1,-1};
        }
        int l=findFirstIdx(nums,target);
        int r=findLastIdx(nums,target);
        return {l, r};
    }
    int findFirstIdx(vector<int>& nums, int target){
        int left=0,right=nums.size()-1;
        while(left<=right){
            int midIdx=left+(right-left)/2;
            if(nums[midIdx]==target){
                // 是不是第一个
                if(midIdx==0||nums[midIdx-1]!=target){
                    return midIdx;
                }else{
                    // 不是第一个 则在左侧
                    right=midIdx-1;
                }
            }else if(nums[midIdx]>target){
                // 更大则在左侧
                right=midIdx-1;
            }else{
                // 更小则在右侧
                left=midIdx+1;
            }
        }
        return -1;
    }
    int findLastIdx(vector<int>& nums, int target){
        int left=0,right=nums.size()-1;
        while(left<=right){
            int midIdx=left+(right-left)/2;
            if(nums[midIdx]==target){
                // 是不是最后一个
                if(midIdx==nums.size()-1||nums[midIdx+1]!=target){
                    return midIdx;
                }else{
                    // 不是最后一个 则在右侧
                    left=midIdx+1;
                }
            }else if(nums[midIdx]>target){
                // 更大则在左侧
                right=midIdx-1;
            }else{
                // 更小则在右侧
                left=midIdx+1;
            }
        }
        return -1;
    }
};
