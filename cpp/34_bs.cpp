class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int l=searchLeftBoundNext(nums,target);
        int r=searchRightBoundNext(nums,target);
        return {l,r};
    }
private:
    // 方法1，[left,right]
    // 寻找左边界的二分搜索
    int searchLeftBound(vector<int>& nums, int target){
        if(nums.size()<1){
            return -1;
        }
        int left=0;
        int right=nums.size()-1;//搜索区间[left,right]，左闭右闭
        while(left<=right){
            // while终止为搜索区间为空，针对搜索区间[left,right]，左闭右闭则是[right+1,right]等于号才满足
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                // 是不是第一个
                if(mid==0||nums[mid-1]!=target){
                    return mid;
                }
                // 不是第一个，由于是寻找左边界，故继续左边找
                right=mid-1;//搜索区间[left,right]，左闭右闭，故right=mid的话[left,mid]会被重复搜索
            }else if(nums[mid]<target){
                // 右边
                left=mid+1;
            }else{
                // 左边
                right=mid-1;//搜索区间[left,right]，左闭右闭，故right=mid的话[left,mid]会被重复搜索
            }
        }
        return -1;
    }
    // 寻找右边界的二分搜索
    int searchRightBound(vector<int>& nums, int target){
        if(nums.size()<1){
            return -1;
        }
        int left=0;
        int right=nums.size()-1;//搜索区间[left,right]，左闭右闭
        while(left<=right){
            // while终止为搜索区间为空，针对搜索区间[left,right]，左闭右闭则是[right+1,right]等于号才满足
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                // 是不是最后一个
                if(mid==nums.size()-1||nums[mid+1]!=target){
                    return mid;
                }
                // 由于是寻找右边界，故继续左边找
                left=mid+1;//搜索区间[left,right]，左闭右闭，故left=mid的话[left,mid]会被重复搜索
            }else if(nums[mid]<target){
                // 右边
                left=mid+1;
            }else{
                // 左边
                right=mid-1;//搜索区间[left,right]，左闭右闭，故right=mid的话[left,mid]会被重复搜索
            }
        }
        return -1;
    }
    // 方法2，[left,right）
        // 寻找左边界的二分搜索
    int searchLeftBoundNext(vector<int>& nums, int target){
        if(nums.size()<1){
            return -1;
        }
        int left=0;
        int right=nums.size();//搜索区间[left,right)，左闭右开
        while(left<right){
            // while终止为搜索区间为空，针对搜索区间[left,right)，左闭右开则是[right,right）小于号才满足
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                // 是不是第一个
                if(mid==0||nums[mid-1]!=target){
                    return mid;
                }
                // 不是第一个，由于是寻找左边界，故继续左边找
                right=mid;//搜索区间[left,right)，左闭右开，故right=mid的话[left,mid]不会被重复搜索
            }else if(nums[mid]<target){
                // 右边
                left=mid+1;
            }else{
                // 左边
                right=mid;//搜索区间[left,right)，左闭右开，故right=mid的话[left,mid]不会被重复搜索
            }
        }
        return -1;
    }
    // 寻找右边界的二分搜索
    int searchRightBoundNext(vector<int>& nums, int target){
        if(nums.size()<1){
            return -1;
        }
        int left=0;
        int right=nums.size();//搜索区间[left,right)，左闭右开
        while(left<right){
            // while终止为搜索区间为空，针对搜索区间[left,right)，左闭右开则是[right,right)小于号才满足
            int mid=left+(right-left)/2;
            if(nums[mid]==target){
                // 是不是最后一个
                if(mid==nums.size()-1||nums[mid+1]!=target){
                    return mid;
                }
                // 由于是寻找右边界，故继续左边找
                left=mid+1;//搜索区间[left,right)，左闭右开，故left=mid的话[left,mid]会被重复搜索
            }else if(nums[mid]<target){
                // 右边
                left=mid+1;
            }else{
                // 左边
                right=mid;//搜索区间[left,right)，左闭右开，故right=mid的话[left,mid）不会被重复搜索
            }
        }
        return -1;
    }
};
