class Solution {
    // 题目：
    // 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
    // 参考：https://blog.csdn.net/qq_51083599/article/details/126634263
public:
    int search(vector<int>& nums, int target) {
        return searchByBoundNext(nums,target);
    }
private:
    int searchByBoundNext(vector<int>& nums, int target){
        int left=0;
        int right=nums.size();//搜索区间[left,right)，左闭右开
        while(left<right){
            // 搜索区间[left,right)，左闭右开，while终止条件是搜索区间为空也就是[right,right)，小于号才满足
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                return mid;
            }else if(nums[mid]<target){
                // 右区间
                left=mid+1;
            }else{
                // 左区间
                right=mid;// 搜索区间[left,right)，左闭右开，mid若作为搜索区间[left,mid)，mid不会被重复搜索
            }
        }
        return -1;
    }
    int searchByBound(vector<int>& nums, int target){
        int left=0;
        int right=nums.size()-1;//搜索区间[left,right]，左闭右闭
        while(left<=right){
            // 搜索区间[left,right]，左闭右闭 while终止条件是搜索区间为空也就是[right+1,right]，等于号才满足left==right+1
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                return mid;
            }else if(nums[mid]<target){
                // 右区间
                left=mid+1;
            }else{
                // 左区间
                right=mid-1;// 搜索区间[left,right]，左闭右闭，mid若作为搜索区间[left,mid]，mid会被重复搜索
            }
        }
        return -1;
    }
};

