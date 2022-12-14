class Solution {
    /*
    给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
    */
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        if(nums.size()<1){
            return {};
        }
        int n=nums.size();
        int finalLen=(1<<n);
        vector<vector<int>> res;

        for(int i=0;i<finalLen;i++){
            vector<int>tmp;
            int bit=i;
            int cnt=0;
            while(bit){
                if((bit&1)==1){
                    tmp.push_back(nums[cnt]);
                }
                bit=(bit>>1);
                cnt++;
            }
            res.push_back(tmp);
        }
        return res;
    }
};
