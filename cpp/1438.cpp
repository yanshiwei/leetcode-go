class Solution {
    /*
    给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit。
    思路：
    利用滑动窗口来寻找符合条件的区间。当区间内绝对差小于等于limit时，则不断向右扩充区间；一旦区间内绝对差大于limit，就从左侧缩小区间直到再次满足条件，然后重复上述过程直到区间末尾。为了计算区间内的最大绝对差，我们要将区间内的数排序，可以通过有序集合来维护区间
    */
public:
    int longestSubarray(vector<int>& nums, int limit) {
        multiset<int>set;// insert erase find count 
        // 底层是红黑树实现的，本身就是一个有序集合，并且允许有相同的元素。
        // 移除元素时要注意是set.erase(set.find(nums[left++]));，而不是set.erase(nums[left++]);，因为集合中可能存在重复元素，前者只移除一个元素，而后者会移除所有该值的元素

        int left=0,right=0;
        int res=0;
        for(;right<nums.size();right++){
            set.insert(nums[right]);
            // 基于multiset有序性排序即可获得最大差
            while(*set.rbegin()-*set.begin()>limit){
                set.erase(set.find(nums[left++]));
            }
            res=max(res,set.size());
        }
        return res;
    }
    int max(int x,int y){
        return x>y?x:y;
    }
};
