class Solution {
    /*
    给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
    */
public:
    int removeDuplicates(vector<int>& nums) {
        if(nums.size()<2){
            return nums.size();
        }
        nums.push_back(10000);//多插入一个使得nums原始数据统一处理
        int j=0;
        for(int i=0;i<nums.size()-1;i++){
            // nums[size-1] is 10000
            if(nums[i]!=nums[i+1]){
                nums[j]=nums[i];
                j++;
            }
        }
        return j;
    }
};
