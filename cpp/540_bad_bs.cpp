class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        //由于给定数组有序 且 常规元素总是两两出现，因此如果不考虑“特殊”的单一元素的话，我们有结论：成对元素中的第一个所对应的下标必然是偶数，成对元素中的第二个所对应的下标必然是奇数。
        // 然后再考虑存在单一元素的情况，假如单一元素所在的下标为 x，那么下标 x 之前（左边）的位置仍满足上述结论，而下标 x之后（右边）的位置由于 xxx 的插入，导致结论翻转。
        // 故存在这样二分型：
        // mid为偶数下标时候，根据上述结论，正常情况下偶数下标的值会与下一值相同，因此如果满足该条件，可以确保 mid之前并没有插入单一元素，此时应该去mid之后区间故left=mid+1；否则right=mid;
        // mid为奇数下标，根据上述结论，正常情况下奇数下标的值会与上一值相同，因此如果满足该条件，可以确保 mid之前并没有插入单一元素，此时应该去mid之后区间故left=mid；否则right=mid-1;
        int n=nums.size();
        int left=0;
        int right=n-1;
        while(left<right){
            // 不同于显式二分，这里不能等于浩，否则陷入循环
            // 因为当left==right时，mid=right，若走到2步骤，right又会等于mid循环了
            int mid=left+(right-left)/2;
            if(mid%2==0){
                // mid为偶数下标时候
                if(mid+1<n&&nums[mid]==nums[mid+1]){
                    // 1.mid之后，且不包括mid，因为mid明确与下一个相等了
                    left=mid+1;
                }else{
                    // 2.mid 之前，且包括mid，因为mid明确与下一个不等，则有可能正好是mid
                    right=mid;
                }
            }else{
                // mid为奇数下标时候
                if(mid-1>=0&&nums[mid-1]==nums[mid]){
                    // 3.mid之后，不包括mid，因为mid明确与上海一个相等了
                    left=mid+1;
                }else{
                    // 4.mid之前，不包括mid，因为mid明确与上一个不等，则单独数字肯定在mid之前，且肯定是偶数下标肯定不是奇数下班
                    right=mid-1;
                }
            }
        }
        return nums[left];
    }
};
