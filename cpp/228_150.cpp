class Solution {
    /*
    i指向每个区间的起始位置，j从i开始向后遍历直到不满足连续递增或者边界，则当前区间结束；
    然后i更新为j+1，作为下一个区间的开始，j继续向后遍历找到下一个区间的结束位置
    */
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector<string>res;
        if(nums.size()<1){
            return res;
        }
        int i=0;
        int j=0;
        while(j<nums.size()){
             // j 向后遍历，直到不满足连续递增(即 nums[j] + 1 != nums[j + 1])
             // 或者 j 达到数组边界，则当前连续递增区间 [i, j] 遍历完毕
             if(j==nums.size()-1||nums[j+1]!=nums[j]+1){
                 string tmp=to_string(nums[i]);
                 if(i!=j){
                     tmp+="->"+to_string(nums[j]);
                 }
                 res.push_back(tmp);
                // i更新为j+1，作为下一个区间的开始
                i=j+1;
             }
             j++;
        }
        return res;
    }
};
