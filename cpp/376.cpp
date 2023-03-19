class Solution {
    /*
    题目：
    如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
    子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度。
    题解，题目要求nums中正数和负数交替的最长子序列的长度，以nums = [1,17,5,10,13,15,10,5,16,8]为例
    i  0 1  2 3  4  5  6  7 8  9
    n  1 17 5 10 13 15 10 5 16 8
    up 1 2  2 4  4  4   4 4  6 6
    do 1 1  3 3  3  3   5 5  5 7
    注意如平峰，则跳过
    https://leetcode.cn/problems/wiggle-subsequence/solutions/518651/376-bai-dong-xu-lie-tan-xin-jing-dian-ti-vyxt/
    */
public:
    int wiggleMaxLength(vector<int>& nums) {
        if(nums.size()<2){
            return nums.size();
        }
        if(nums.size()==2&&nums[0]!=nums[1]){
            return 2;
        }

        int up=1;// 定义变量up 表示当前递增的最长摆动序列,初始值为1.
        int down=1;// 定义变量down表示当前递减的最长摆动序列,初始值为1.
        for(int i=1;i<nums.size();i++){
            if(nums[i]>nums[i-1]){
                // 当前递增
                up=down+1;
            }else if(nums[i]<nums[i-1]){
                // 当前递减
                down=up+1;
            }else{
                // 平峰
                continue;
            }
        }
        return max(up,down);
    }
};
