class Solution {
    /*
    给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。
    */
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target){
        if(nums.size()<4){
            return {};
        }
        vector<vector<int>> res;
        sort(nums.begin(),nums.end());
        int n=nums.size();
        // 先固定nums[i]
        for(int i=0;i<n-3;i++){
            // 第一个的起始位置i，故右边必须至少3个
            if(i>0&&nums[i]==nums[i-1]){
                continue;//避免重复
            }
            // 再固定nums[j]
            for(int j=i+1;j<n-2;j++){
                //第二个的起始位置j，故右边必须至少2个
                if(j>i+1&&nums[j]==nums[j-1]){
                    continue;
                }
                // 开启相向双指针
                int left=j+1;
                int right=n-1;
                while(left<right){
                    // 不可以重复使用同一个元素故这里left<right
                    long sum=(long)nums[i]+nums[j]+nums[left]+nums[right];//long 避免溢出
                    if(sum==target){
                        res.push_back({nums[i],nums[j],nums[left],nums[right]});
                        // 判断左界和右界是否和下一位置重复，去除重复解。并同时将L,R 移到下一位置，寻找新的解
                        while(left<right&&nums[left]==nums[left+1]){
                            left++;
                        }
                        while(left<right&&nums[right]==nums[right-1]){
                            right--;
                        }          
                        // 当前left/right位置是这个值的最后一位  
                        left++;
                        right--;       
                    }else if(sum>target){
                        // 过大
                        right--;
                    }else{
                        // 过小
                        left++;
                    }
                }
            }
        }
        return res;
    }
/*
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> quadruplets;
        if (nums.size() < 4) {
            return quadruplets;
        }
        sort(nums.begin(), nums.end());//方便双向指针且方便去重
        int length = nums.size();
        // 先固定nums[i]
        for (int i = 0; i < length - 3; i++) {
            // 第一个的起始位置i，故右边必须至少3个
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;// 避免重复
            }
            // 再固定nums[j]
            for (int j = i + 1; j < length - 2; j++) {
                //第二个的起始位置j，故右边必须至少2个
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    continue;
                }
                // 开启相向双指针
                int left = j + 1, right = length - 1;
                while (left < right) {
                    long sum = (long) nums[i] + nums[j] + nums[left] + nums[right];
                    if (sum == target) {
                        quadruplets.push_back({nums[i], nums[j], nums[left], nums[right]});
                        // 然后将左指针右移直到遇到不同的数，将右指针左移直到遇到不同的数避免重复
                        while (left < right && nums[left] == nums[left + 1]) {
                            left++;
                        }
                        left++;// skip the same value to first different value nums[left+1]
                        while (left < right && nums[right] == nums[right - 1]) {
                            right--;
                        }
                        right--;
                    } else if (sum < target) {
                        left++;
                    } else {
                        right--;
                    }
                }
            }
        }
        return quadruplets;
    }
    */
};

