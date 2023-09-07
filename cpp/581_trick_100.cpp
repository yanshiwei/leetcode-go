class Solution {
    /*
    题目：
    给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的 最短 子数组，并输出它的长度。
示例 1：
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
    题解：
    https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/solutions/911869/581-zui-duan-wu-xu-lian-xu-zi-shu-zu-si-8rivt/
    数组实际上可以分为三段，有序A+无序B+有序C，且两段有序肯定有前面一段都小于后面一段的，只需要确定无序的边界就可以。
     做法就是维护一个max一个min来帮助确定边界， 先看max，一边遍历一边更新当前最大值，若便利时候发现比当前最大值小，说明就不是有序的，而是无序的，这时候就是无序的一个临时边界，一直便利到结束，最后被更新到的就是最终的右边界。
     min类似，从右到左遍历，正常应该是逐渐变小，如果存在更大的，说明是无序序列的一个临时左边界。
    */
public:
    const int int_max=0x7fffffff;
    const int int_min=0x80000000;
    int findUnsortedSubarray(vector<int>& nums) {
        int n=nums.size();
        if(n<1){
            return 0;
        }
        int maxv=int_min;
        int minv=int_max;
        int left=-1;
        int right=-1;
        for(int i=0;i<n;i++){
            if(nums[i]<maxv){
                // 从左到右遍历，发现更小的，说明就是一个临时右边界
                right=i;
            }else{
                // 更大的直接更新maxv
                maxv=max(maxv,nums[i]);
            }
            if(nums[n-1-i]>minv){
                // 从右到左遍历，发现更大的，说明就是一个临时左边界
                left=n-1-i;
            }else{
                // 更大的直接更新maxv
                minv=min(minv,nums[n-1-i]);
            }            
        }
        return right==-1?0:(right-left+1);
    }
};
