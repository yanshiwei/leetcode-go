class Solution {
    /*
    给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）

    map:insert erase find count lower_bound size empty
    */
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        unordered_map<int,int>set;//由于同一个数字在两个数组中都可能出现多次，因此需要用哈希表存储每个数字出现的次数。对于一个数字，其在交集中出现的次数等于该数字在两个数组中出现次数的最小值。
        for(int i=0;i<nums1.size();i++){
            set[nums1[i]]++;
        }
        vector<int>res;
        for(int i=0;i<nums2.size();i++){
            if(set.count(nums2[i])){
                res.push_back(nums2[i]);
                set[nums2[i]]--;
                if(set[nums2[i]]==0){
                    //出现的次数等于该数字在两个数组中出现次数的最小值
                    set.erase(nums2[i]);
                }
            }
        }
        return res;
    }
};
