class Solution {
    /*
    题目：
    峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
    题解：
    基于，nums[-1] = nums[n] = -∞证明数组必有峰值：
    1.1 如果数组长度为1，由于nums[-1] = nums[n] = -∞则nums[0]就是峰值
    1.2 如果数字长度大于1，nums[-1]= -∞，则有两种情况
        1.2.1 如果nums[0]>nums[1]，那么最左边元素 nums[0]就是峰值（结合左边界为负无穷）；
        1.2.2 如果nums[0]<nums[1]，由于已经存在明确的nums[0] 和 nums[1]大小关系，我们将 nums[0]看做边界， nums[1]看做新的最左侧元素，继续往右进行分析：
            1.2.2.1 到达数组末尾前，出现 nums[i]>nums[i+1]，说明存在峰值位置 i；
            1.2.2.2 到达数组最右侧，还没出现nums[i]>nums[i+1]，说明数组严格递增。此时结合右边界可以看做负无穷，可判定nums[n−1] 为峰值
    总之必有峰值
    然后，我们证明二分可以找到峰值，注意：这里的「二段性」其实是指：在以 mid为分割点的数组上，根据 nums[mid] 与 nums[mid±1] 的大小关系，可以确定其中一段满足「必然有解」，另外一段不满足
    对于一个满足nums[x]>nums[x+1] 的位置：
    则x的y右侧一定存在峰值而非左侧；
    反之依然
    */
public:
    int findPeakElement(vector<int>& nums) {
        int left=0;
        int right=nums.size()-1;
        while(left<right){
            int mid=left+(right-left)/2;
            if(nums[mid]>nums[mid+1]){
                //左侧
                right=mid;
            }else{
                // 右侧
                left=mid+1;
            }
        }
        return left;
    }
};
