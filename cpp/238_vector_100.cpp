class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        // 计算nums[i]之前乘积
        int k=1;
        vector<int>res(nums.size(),0);
        for(int i=0;i<nums.size();i++){
            res[i]=k;
            k*=nums[i];
        }
        // 计算nums[i]之后乘积
        k=1;
        for(int i=nums.size()-1;i>=0;i--){
            res[i]*=k;
            k*=nums[i];
        }
        return res;
    }
};
