class Solution {
    /*
    给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法且不使用额外空间来解决此问题。
    */
public:
    int singleNumber(vector<int>& nums) {
        int res=0;
        for(int i=0;i<32;i++){
            int cnt=0;
            for(auto&num:nums){
                cnt+=(num>>i)&1;
            }
            if(cnt%3==1){
                res|=(1<<i);
            }
        }
        return res;
    }
};
