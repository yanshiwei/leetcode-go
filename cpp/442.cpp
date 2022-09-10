class Solution {
    /*
    给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
    */
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int>res;
        int nLen=nums.size();
        if(nLen<2){
            return res;
        }
        for(int i=0;i<nums.size();i++){
            if(nums[i]<1||nums[i]>nLen){
                return res;
            }
        }
        for(int i=0;i<nums.size();i++){
            while(nums[i]!=nums[nums[i]-1]){
                int tmp=nums[i];
                nums[i]=nums[tmp-1];
                nums[tmp-1]=tmp;
            }
        }
        for(int i=0;i<nums.size();i++){
            if(nums[i]!=i+1){
                res.push_back(nums[i]);
            }
        }
        return res;
    }
};
