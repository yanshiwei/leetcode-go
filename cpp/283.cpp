class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        // 同向双指针
        // 左指针指向当前已经处理好的序列的下一个，右指针指向待处理序列的头部
        // 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
        // 注意：1 左指针左边均为非零数； 2 右指针左边直到左指针处均为零(左指针可能为0)
        int n=nums.size();
        int left=0;//当前已经处理好的序列的下一个 左指针左边均为非零数 (左指针可能为0)
        int right=0;//右指针指向待处理序列的头部 右指针左边直到左指针处均为零
        /*
        题目：
        给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
        思路：
        |---非0部分---|left|right---未处理部分---|
        */
        while(right<n){
            if(nums[right]){
                swap(nums[left],nums[right]);
                left++;
            }
            right++;
        }
    }
};
