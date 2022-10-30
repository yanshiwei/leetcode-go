class Solution {
    /*
    给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。=
如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
    */
public:
    int minOperations(vector<int>& nums, int x) {
        // 将题意转为 最长连续子数组 nums[left, right]，使得其和为sum-x，可运用滑动窗口进行求解
        long target=sum(nums)-x;
        if(target<0){
            return -1;//1 <= nums[i]
        }
        if(target==0){
            return nums.size();
        }
        int left=0;
        int right=0;
        int res=-1;
        long total=0;//滑动窗口中的数字总和
        while(right<nums.size()){
            total+=nums[right];
            while(total>target){
                total-=nums[left];
                left++;
            }
            if(total==target){
                res=max(res,right-left+1);
            }
            right++;
        }
        if(res==-1){
            return -1;
        }
        return nums.size()-res;
    }
private:
    long sum(vector<int>& nums){
        long res=0;
        for(int i=0;i<nums.size();i++){
            res+=nums[i];
        }
        return res;
    }
};
