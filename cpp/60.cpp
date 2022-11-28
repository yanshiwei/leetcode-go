class Solution {
    /*
    题目：
    给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。
题解：
    https://leetcode.wang/leetCode-60-Permutation-Sequence.html
    */
public:
int fac(int n){
    if(n<=1){
        return 1;
    }
    return n*fac(n-1);
}
    string helper(vector<int>&nums,int n,int k){
        if(n==1){
            //把剩下的最后一个数字返回就可以了
            return to_string(nums[0]);
        }
        int perGroupNum=fac(n-1);//每组的个数
        int groupNo=(k-1)/perGroupNum;//第几组
        int num=nums[groupNo];
        nums.erase(nums.begin()+groupNo);
        k=k%perGroupNum;//更新下次的 k 
        k=k==0?perGroupNum:k;
        return to_string(num)+helper(nums,n-1,k);
    }
    string getPermutation(int n, int k) {
        if(n<1){
            return "";
        }
        vector<int>nums;
        for(int i=0;i<n;i++){
            nums.push_back(i+1);
        }
        return helper(nums,n,k);
    }
};
