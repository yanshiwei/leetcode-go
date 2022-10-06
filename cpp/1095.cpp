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
