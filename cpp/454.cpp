class Solution {
    /*
    给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/
public:
    int fourSumCount(vector<int>& nums1, vector<int>& nums2, vector<int>& nums3, vector<int>& nums4) {
        //A 和 B，我们使用二重循环对它们进行遍历，得到所有A[i]+B[j] 的值并存入哈希映射中。对于哈希映射中的每个键值对，每个键表示一种A[i]+B[j]，对应的值为 A[i]+B[j] 出现的次数
        unordered_map<int,int>hash;// insert erase find count
        for(auto &a:nums1){
            for(auto &b:nums2){
                hash[a+b]++;
            }
        }
        int res=0;
        for(auto &c:nums3){
            for(auto &d:nums4){
                if(hash.count(-c-d)){
                    res+=hash[-c-d];
                }
            }
        }        
        return res;
    }
};
