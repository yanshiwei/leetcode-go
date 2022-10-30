class Solution {
    /*
 把「最多可以把 K 个 0 变成 1，求仅包含 1 的最长子数组的长度」转换为 「找出一个最长的子数组，该子数组内最多允许有 K 个 0 」
 代码思路：
 1. 使用 left和 right两个指针，分别指向滑动窗口的左右边界
 2. right 主动右移逻辑：right指针每次移动一步。当 A[right]为 0，说明滑动窗口内增加了一个 0；
 3. left 被动右移：判断此时窗口内 0的个数，如果超过了 K，则 left指针被迫右移，直至窗口内的 0的个数小于等于 K 为止
 4. 滑动窗口长度的最大值就是所求
    */
public:
    int longestOnes(vector<int>& nums, int k) {
        //可知本题是求最大连续子区间，可以使用滑动窗口方法。滑动窗口的限制条件是：窗口内最多有 K 个 0
        int maxLen=0;
        int zeroCnt=0;
        int left=0;
        int right=0;
        while(right<nums.size()){
            if(nums[right]==0){
                zeroCnt++;
            }
            while(zeroCnt>k){
                if(nums[left]==0){
                    zeroCnt--;
                }
                left++;
            }
            maxLen=max(maxLen,right-left+1);
            right++;
        }
        return maxLen;
    }
};
