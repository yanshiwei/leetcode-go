// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
    //[good,good,bad,bad]，二分性质，二分算法
public:
    int firstBadVersion(int n) {
       int left=1;//版本最小
       int right=n;//版本最大
       while(left<right){
           int mid=left+(right-left)/2;
           if(isBadVersion(mid)){
               // 则首个错误版本一定在闭区间 [left, mid]
               right=mid;
           }else{
               // 首个错误版本一定在闭区间 [m + 1, right]
               left=mid+1;
           }
       }
       return left;
    }
};
