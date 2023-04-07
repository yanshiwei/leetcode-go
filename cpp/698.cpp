class Solution {
    /*
    给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
    回溯：https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/solutions/1441006/by-lfool-d9o7/?orderBy=most_votes
    */
public:
    bool dfs(vector<int>&nums,int idx,vector<int>&bucket,int k,int target){
        // 结束条件：已经处理完所有球
        if(idx==nums.size()){
        // 因为有bucket[i]+nums[idx]>target条件判断，且和一定为sum，而每一个都不大于sum/k
        // 已知a+b+c=3m且a<=m,b<=m,c<=m ，有a=3m-(b+c),b+c<=2m,所以a=3m-(b+c)>=m ,因此m<=a<=m
        // 所以 index == num.length 时，所有球已经按要求装入所有桶，所以肯定是一个满足要求的解
            return true;
        }
        // nums[index] 开始做选择,选择某个桶
        for(int i=0;i<k;i++){
            // 如果当前桶和上一个桶内的元素和相等，那么 nums[index] 选择上一个桶和选择当前桶可以得到的结果是一致的
            if(i>0&&bucket[i]==bucket[i-1]){
                continue;
            }
            // 剪枝：放入球后超过 target 的值，选择下一个桶
            if(bucket[i]+nums[idx]>target){
                continue;
            }
            // 做选择：放入 i 号桶
            bucket[i]+=nums[idx];
            // 处理下一个球，即 nums[index + 1]
            if(dfs(nums,idx+1,bucket,k,target)){
                return true;
            }
            // 撤销选择：挪出 i 号桶
            bucket[i]-=nums[idx];
        }
        // k 个桶都不满足要求
        return false;
    }
    // bool dfs(vector<int>&nums,int idx,vector<int>&bucket,int k,int target){
    //     // 结束条件：已经处理完所有球
    //     if(idx==nums.size()){
    //         // 判断现在桶中的球是否符合要求 -> 路径是否满足要求
    //         for(int i=0;i<k;i++){
    //             if(bucket[i]!=target){
    //                 return false;
    //             }
    //         }
    //         return true;
    //     }
    //     // nums[index] 开始做选择,选择某个桶
    //     for(int i=0;i<k;i++){
    //         // 剪枝：放入球后超过 target 的值，选择下一个桶
    //         if(bucket[i]+nums[idx]>target){
    //             continue;
    //         }
    //         // 做选择：放入 i 号桶
    //         bucket[i]+=nums[idx];
    //         // 处理下一个球，即 nums[index + 1]
    //         if(dfs(nums,idx+1,bucket,k,target)){
    //             return true;
    //         }
    //         // 撤销选择：挪出 i 号桶
    //         bucket[i]-=nums[idx];
    //     }
    //     // k 个桶都不满足要求
    //     return false;
    // }
    bool canPartitionKSubsets(vector<int>& nums, int k) {
        int sum=accumulate(nums.begin(),nums.end(),0);
        if(sum%k!=0){
            // 不能平分
            return false;
        }
        int target=sum/k;
        vector<int>bucket(k+1,0);//每个桶的和
        sort(nums.rbegin(),nums.rend());//反向排序,先让值大的元素选择桶，这样可以增加剪枝的命中率，从而降低回溯的概率
        return dfs(nums,0,bucket,k,target);
    }
};
