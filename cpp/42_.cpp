class Solution {
    /*
    双指针
    每一个柱子所接雨水和即为所接雨水
    对于柱子i，其所接雨水为，其左侧最高和右侧最高的较低值减去本柱子高度再乘以宽度（宽带为1），当然此时本柱子高度大于这个较低值时候就接不了雨水：
    left_max<right_max：
        res+=max(left_max-height[left],0)
        left_max=max(left_max,height[left])
    反之亦然
    */
public:
    int trap(vector<int>& height){
        if(height.size()<1){
            return 0;
        }
        int res=0;
        int left=0;
        int right=height.size()-1;
        int left_max=0;
        int right_max=0;
        while(left<=right){
            //[left,right]闭区间所以等于号
            if(left_max<right_max){
                res+=max(left_max-height[left],0);
                left_max=max(left_max,height[left]);
                left++;
            }else{
                res+=max(right_max-height[right],0);
                right_max=max(right_max,height[right]);
                right--;               
            }
        }
        return res;
    }
    // int trap(vector<int>& height) {
    //     //给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
    //     if(height.size()<1){
    //         return 0;
    //     }
    //     int left=0,right=height.size()-1;
    //     int left_max=0,right_max=0;
    //     int res=0;
    //     while(left<=right){
    //         if(height[left]<height[right]){
    //             // left_max height[left]更新 故left_max<right_max 右侧构成一个墙壁
    //             if(height[left]<left_max){
    //                 // 可以蓄水
    //                 res+=left_max-height[left];
    //             }else{
    //                 left_max=height[left];
    //             }
    //             left++;
    //         }else{
    //             // right_max 由height[right]更新 故left_max>right_max 左侧构成一个墙壁
    //             if(height[right]<right_max){
    //                 res+=right_max-height[right];
    //             }else{
    //                 right_max=height[right];
    //             }
    //             right--;
    //         }
    //     }
    //     return res;
    // }
};
