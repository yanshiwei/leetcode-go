class Solution {
    // 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
    // 下一个更大值，单调递增栈
    // 循环数组，要找的目标，要么在当前元素之前，要么在之后，因此两次遍历一定能覆盖到
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        int n=nums.size();
        if(n<1){
            return {};
        }
        vector<int>res(n,-1);//ret提前初始化好，这样每个位置通过赋值方式避免2次数组处理时有重复
        stack<int>st;//最大值找不到默认-1，故需不需要st.push(max)来处理
        for(int i=0;i<2*n-1;i++){
            if(st.empty()||nums[st.top()]>=nums[i%n]){
                st.push(i%n);
            }else{
                while(!st.empty()&&nums[st.top()]<nums[i%n]){
                    res[st.top()]=nums[i%n];
                    st.pop();
                }
                st.push(i%n);
            }
        }
        return res;
    }
};
