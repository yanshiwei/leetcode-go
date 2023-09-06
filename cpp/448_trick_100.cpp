class Solution {
    /*
    1<=nums[i]<=n
    如在(nums[i]-1)位置处+n，则该位置值肯定大于n
    注意这里要取模，因为该位置可能已经之前加过n，早就超过n
    执行一遍后，查看哪个位置值<=n，如有则就是缺少的值
    */
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        if(nums.size()<1){
            return {};
        }
        vector<int>res;
        int n =nums.size();
        for(int i=0;i<n;i++){
            nums[(nums[i]-1)%n]+=n;
        }
        for(int i=0;i<n;i++){
            if(nums[i]<=n){
                res.push_back(i+1);
            }
        }
        return res;
    }
};
