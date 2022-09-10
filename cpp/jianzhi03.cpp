class Solution {
    /*
    在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
    */
public:
    int findRepeatNumber(vector<int>& nums) {
        int nLen=nums.size();
        if(nLen<2){
            return -1;
        }
        for(int i=0;i<nLen;i++){
            if(nums[i]<0||nums[i]>nLen-1){
                return -1;
            }
        }
        int dup=-1;
        int find=false;
        for(int i=0;i<nLen;i++){
            while(nums[i]!=i){
                if(nums[i]==nums[nums[i]]){
                    dup=nums[i];
                    find=true;
                    break;
                }
                int tmp=nums[i];
                nums[i]=nums[tmp];
                nums[tmp]=tmp;
            }
            if(find){
                break;
            }
        }
        return dup;
    }
};
