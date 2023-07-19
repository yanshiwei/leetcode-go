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
    1 先找到山顶元素 mountaintop 所在的下标
    2 在前有序且升序数组中找 target 所在的下标，如果找到了，就返回，如果没有找到，才执行第 3 步
    3 如果步骤 2 找不到，就在后有序且降序数组中找 target 所在的下标
    */
public:
    int findInMountainArray(int target, MountainArray &mountainArr) {
        // 1 找到峰值
        int left=0;
        int right=mountainArr.length()-1;//[left,right]
        while(left<=right){
            int mid=left+(right-left)/2;
            if(mountainArr.get(mid)>mountainArr.get(mid+1)){
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        int big=left;
        // 2 find in left array, the array is sorted
        left=0;
        right=big;
        int res=-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if(mountainArr.get(mid)==target){
                res=mid;
                break;
            }
            if(mountainArr.get(mid)>target){
                // left
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        if(res!=-1){
            return res;
        }
        // find in right array, the array is sorted
        left=big;
        right=mountainArr.length()-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if(mountainArr.get(mid)==target){
                res=mid;
                break;
            }
            if(mountainArr.get(mid)>target){
                // right for desc
                left=mid+1;
            }else{
                right=mid-1;
            }            
        }
        return res;
    }
};
