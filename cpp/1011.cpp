class Solution {
    /*
    传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
    */
public:
    int shipWithinDays(vector<int>& weights, int days) {
        // 船的最大载重的最小化 二分
        // 1. 范围
        int left=0;
        // 最小要一天运一个，此时max of arr
        for(int i=0;i<weights.size();i++){
            if(left<weights[i]){
                left=weights[i];
            }
        }
        // 最大一天全运完，此时sum of arr
        int right=0;
        for(int i=0;i<weights.size();i++){
            right+=weights[i];
        }
        while(left<right){
            int mid=left+(right-left)/2;
            if(check(weights,mid,days)){
                // 船的最大载重mid时days天可完成 为了最小化
                right=mid;
            }else{
                left=mid+1;
            }
        }
        return left;
    }
private:
// 船的最大载重mid时是否mid天内完成
    bool check(vector<int>& weights,int mid,int days){
        int res=1;
        int base=0;
        for(auto w:weights){
            if(base+w>mid){
                res++;// 新的一天
                base=w;
            }else{
                base+=w;
            }
        }
        return res<=days;
    }
};
