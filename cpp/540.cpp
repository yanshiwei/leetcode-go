class Solution {
    /*
    题目：
    给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
请你找出并返回只出现一次的那个数。
你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
    思路：
    如果没有缺失，则
    nums 1 1 2 2 3 3 4 4 8 8
    idx: 0 1 2 3 4 5 6 7 8 9
    即连续一对偶数下标/奇数下标结果值相同
    若有缺失
    nums 1 1 2 3 3 4 4 8 8
    idx: 0 1 2 3 4 5 6 7 8
    有了缺失后，缺失的左侧仍然是偶数奇数对，右侧则是先奇后偶。
    如果 mid 所在的[偶数下标, 奇数下标]的值相等，说明前面半段没有缺失的数，那么，缺失的数肯定在后半段，反之，则在前半段。
    */
public:
    int singleNonDuplicate(vector<int>& nums) {
        int left=0,right=nums.size()-1;
        //这里如果判断条件写 left <= right 的话，那么在 left == right 的条件下进入循环时，若数组长度只有 1 ，此时 left == right == mid == 1，计算 nums[mid + 1] 会导致越界
        while(left<right){
            int mid=left+(right-left)/2;
            if((mid&1)==1){
                //如果mid是奇数则将mid减1 保证mid是偶数
                mid--;
            }
            if(nums[mid]==nums[mid+1]){
                // in the right
                left=mid+2;//由于是对偶数位进行二分查找，移动左边界时要加 2，而不是加 1
            }else{
                // in the left include mid
                right=mid;
            }
        }
        return nums[left];
    }
};
