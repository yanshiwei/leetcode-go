class Solution {
public:
/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

*/
    int removeElement(vector<int>& nums, int val) {
        int cnt=0;//当前位置前面有几个与val值相同的元素
        for(int i=0;i<nums.size();i++){
            if (nums[i]==val){
                cnt++;
            }else{
                nums[i-cnt]=nums[i];//当前位置元素需要向前移动多少位
            }
        }
        return nums.size()-cnt;
    }
};
