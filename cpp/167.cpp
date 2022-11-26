class Solution {
    /*
    给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
    */
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        // 相向双指针
        int left=0,right=numbers.size()-1;
        while(left<right){
            int sum=numbers[left]+numbers[right];
            if(sum<target){
                // 过小，增加
                left++;
            }else if(sum>target){
                // 过大，减小
                right--;
            }else{
                return {left+1,right+1};//下标从 1 开始
            }
        }
        return {-1,-1};
    }
};