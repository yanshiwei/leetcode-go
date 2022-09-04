class Solution {
public:
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

*/
    void sortColors(vector<int>& nums) {
         // 维护三个区域，小于1的区域，等于1的区域，大于1的区域
        //双指针，left,right指向边界，cur用来遍历 left指向0，right指向2
        int left=-1,right=nums.size();
        int cur=0;
        while(cur<right){
            if (nums[cur]<1){
                // 小于1，与left的右一个数交换
                swap(nums[cur],nums[left+1]);
                cur++;
                left++;
            }else if(nums[cur]>1){
                // 大于1，与right的左一个数交换，i不变(因为交换过来的数还没看)
                // cur不动，因为从右边换过来的数还不知道和1比谁大
                swap(nums[cur],nums[right-1]);
                right--;
            }else{
                cur++;
            }
        }
    }
};
