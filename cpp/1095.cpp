/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * class MountainArray {
 *   public:
 *     int get(int index);
 *     int length();
 * };
 */

class Solution {
    /*
    在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]
你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：
MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
MountainArray.length() - 会返回该数组的长度
    */
public:
    int findInMountainArray(int target, MountainArray &mountainArr) {
        // 1 find 峰值
        int left=0,right=mountainArr.length()-1;
        while(left<right){
            int mid=left+(right-left)/2;
            if(mountainArr.get(mid)>mountainArr.get(mid+1)){
                // 因为山峰数组故mid肯定走不到length-1
                right=mid;
            }else{
                left=mid+1;
            }
        }
        int big=left;
        // 2 find in left is up 
        left=0,right=big;
        int findRes=-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if (mountainArr.get(mid)==target){
                findRes=mid;
                break;
            }
            if(mountainArr.get(mid)>target){
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        if(findRes>-1){
            return findRes;
        }
        // 3 find in right is down
        left=big,right=mountainArr.length()-1;
        findRes=-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if (mountainArr.get(mid)==target){
                findRes=mid;
                break;
            }
            if(mountainArr.get(mid)>target){
                left=mid+1; 
            }else{
                right=mid-1;
            }            
        }
        if(findRes>-1){
            return findRes;
        }
        return -1;
    }
};
