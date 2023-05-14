class Solution {
    /*
    题目：
    给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
    题解：
    [5,5,1,2,2,3,2]
    101
    101
    001
    010
    010
    011
    010
    每一列相加，出现这样的数的对应的数字的数目肯定也都大于n/2
    */
public:
    int majorityElement(vector<int>& nums) {
        int res=0;
        for(int i=0;i<32;i++){
            int mask=(1<<i);
            int cnt=0;
            for(int j=0;j<nums.size();j++){
                if((mask&nums[j])==mask){
                    cnt++;
                }
            }
            if(cnt>nums.size()/2){
                res|=mask;
            }
        }
        return res;
    }
};
