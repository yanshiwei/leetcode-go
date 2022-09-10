
class Solution {
    /*
    给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
    */
public:
int countRange(vector<int>& nums,int left,int right){
    int res=0;
    for(int i=0;i<nums.size();i++){
        if(nums[i]>=left&&nums[i]<=right){
            res++;
        }
    }
    return res;
}
    int findDuplicate(vector<int>& nums) {
        int nLen=nums.size();
        if(nLen<2){
            return 0;
        }
        for(int i=0;i<nLen;i++){
            if(nums[i]<1||nums[i]>nLen-1){
                return 0;
            }
        }
        int start=1;
        int end=nLen-1;
        // half search
        while(end>=start){
            int mid=((end - start) >> 1) + start;
            int cnt=countRange(nums,start,mid);
            if(start==end){
                // 一个数
                if(cnt>1){
                    return start;
                }
                break;
            }
            if(cnt>(mid-start+1)){
                // 重复的在start-mid
                end=mid;
            }else{
                // 重复的在mid+1-end
                start=mid+1;
            }
        }
        return 0;
    }
};
