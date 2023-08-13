class Solution {
    /*
    给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
    */
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int>hash(nums.begin(),nums.end());
        int res=0;
        for(int i=0;i<nums.size();i++){
            int cur=nums[i];
            // 如果当前数cur前驱cur-1不存在，则以cur开始向后尝试寻找连续数并时刻更新最大值
            if(hash.count(cur-1)==0){
                int x=cur+1;
                while(hash.count(x)){
                    x++;
                }
                res=max(res,x-cur);
            }
        }
        return res;
    }
};
