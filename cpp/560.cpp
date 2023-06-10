class Solution {
    // 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
    // 前缀和+hash
    // 假如存在区间[left,right]，使得在[left,right]这个区间的子数组的和为k。换句话说，就是前right项和减去前left-1项和等于k，即前left-1项和等于前right项和减去k。
    // 也就是在扫描数组的同时，假设当前扫到第right位，记录它的前right项和sum，用该和减去k，即sum-k，判断sum-k是否为之前出现过也就是前left-1项和
public:
    int subarraySum(vector<int>& nums, int k) {
        int n=nums.size();
        vector<int>preSum(n+1,0);
        // 预处理，得到前缀和，presum[0]=0;presum[1]=presum[0]+nums[0](前缀和数组下标默认从 1 开始)
        for(int i=1;i<=n;i++){
            preSum[i]=preSum[i-1]+nums[i-1];
        }
        unordered_map<int,int>hash;// key is presum ,value is fre
        hash.insert(make_pair(0,1));// presum[0]
        int res=0;
        for(int i=1;i<=n;i++){
            int sum=preSum[i];
            int d=sum-k;
            if(hash.count(d)>0){
                res+=hash[d];
            }
            if(hash.count(sum)>0){
                hash[sum]+=1;
            }else{
                hash[sum]=1;
            }
        }
        return res;
    }
};
