class Solution {
    /*
    nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
    */
public:
    vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
        // 先找num2的下一个最大值，递增栈
        stack<int>st;
        unordered_map<int,int>table;//key is num2[i], value is idx of bigger
        for(int i=0;i<nums2.size();i++){
            if(st.empty()||nums2[st.top()]>nums2[i]){
                st.push(i);
            }else{
                while(!st.empty()&&nums2[st.top()]<nums2[i]){
                    int idx=st.top();
                    st.pop();
                    table[nums2[idx]]=i;
                }
                st.push(i);
            }
        }
        vector<int>res(nums1.size(),-1);
        for(int i=0;i<nums1.size();i++){
            if(table.count(nums1[i])>0){
                res[i]=nums2[table[nums1[i]]];
            }
        }
        return res;
    }
};
