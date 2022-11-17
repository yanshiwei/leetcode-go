class Solution {
    /*给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
*/
public:
    bool canJump(vector<int>& nums) {
        if(nums.size()<1){
            return false;
        }
        //为当前能够到达的最大位置 
        int maxDis=0;
        for(int i=0;i<nums.size();i++){
            //遍历元素位置下标大于当前能够到达的最大位置下标，不能到达
            if(i>maxDis){
                return false;
            }
            //表示当前元素能跳的最大覆盖范围，每次我都只往右跳一格
            maxDis=max(maxDis,i+nums[i]);
        }
        return true;
    }
};
