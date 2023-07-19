class Solution {
    /*
    给定一个  无重复元素 的 有序 整数数组 nums 。
返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
列表中的每个区间范围 [a,b] 应该按如下格式输出：
"a->b" ，如果 a != b
"a" ，如果 a == b
    */
public:
    vector<string> summaryRanges(vector<int>& nums) {
        if(nums.size()<1){
            return {};
        }
        vector<string>res;
        int start=nums[0];
        int end=nums[0];
        int n=nums.size();
        for(int i=0;i<n-1;i++){
            if(nums[i+1]!=nums[i]+1){
                //不连续
                end=nums[i];
                if(start!=end){
                    res.push_back(to_string(start)+"->"+to_string(end));
                }else{
                    res.push_back(to_string(start));
                }
                start=nums[i+1];
            }
        }
        end=nums[n-1];
        if(start!=end){
            res.push_back(to_string(start)+"->"+to_string(end));
        }else{
            res.push_back(to_string(start));
        } 
        return res;
    }
};
