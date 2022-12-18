class Solution {
    /*
    因为给定数组是有序的，所以相同元素必然连续。我们可以使用双指针解决本题，遍历数组检查每一个元素是否应该被保留，如果应该被保留，就将其移动到指定位置。具体地，我们定义两个指针 slow 和 fast分别为慢指针和快指针，其中慢指针表示处理出的数组的长度，快指针表示已经检查过的数组的长度，即 nums[fast] 表示待检查的第一个元素，nums[slow−1]为上一个应该被保留的元素所移动到的指定位置。
因为本题要求相同元素最多出现两次而非一次，所以我们需要检查上上个应该被保留的元素 nums[slow−2]是否和当前待检查元素 nums[fast]相同。当且仅当 nums[slow−2]=nums[fast]时，当前待检查元素 nums[fast]\textit{nums}[\textit{fast}]nums[fast] 不应该被保留。最后，slow 即为处理好的数组的长度。
特别地，数组的前两个数必然可以被保留，因此对于长度不超过 2 的数组，我们无需进行任何处理，对于长度超过 2 的数组，我们直接将双指针的初始值设为 2 即可。

    */
public:
    int removeDuplicates(vector<int>& nums) {
        if(nums.size()<3){
            return nums.size();
        }
        int slow=2;
        int fast=2;
        while(fast<nums.size()){
            if(nums[slow-2]!=nums[fast]){
                // 当前2位是11 第三位是2时
                // 当前2位是12 第三位是2时
                // 都需要保留
                nums[slow]=nums[fast];
                slow++;
            }
            fast++;
        }
        return slow;
    }
};
