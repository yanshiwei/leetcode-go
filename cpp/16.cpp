class Solution {
    /*
    给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解
    */
public:
    int abs(int a){
        return a>0 ? a:-a ;
    }
    // 在数组 nums 中，进行遍历，每遍历一个值利用其下标i，形成一个固定值 nums[i]
    int threeSumClosest(vector<int>& nums, int target) {
        // init ans with head 3, 升序数组中，三元组的元素之和的最小值head 3 之和 
        int ans=nums[0]+nums[1]+nums[2];
        int n=nums.size();
        sort(nums.begin(),nums.end());
        for(int i=0;i<n;i++){
            // 相向双指针
            int left = i+1,right=nums.size()-1;
            while(left<right){
                int sum= nums[left]+nums[right]+nums[i];
                // 根据 sum = nums[i] + nums[start] + nums[end] 的结果，判断 sum 与目标 target 的距离，如果更近则更新结果 ans
                if(abs(sum-target)<abs(target-ans))
                ans=sum;
                if(sum<target)
                left++;
                else if(sum>target)
                right--;
                else
                return ans;
            }
        }
             return ans;          
    }
};
