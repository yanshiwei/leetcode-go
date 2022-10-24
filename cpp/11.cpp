class Solution {
    /*
    给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量
    */
public:
    int maxArea(vector<int>& height) {
        int res=0;
        // 相向双指针
        int left=0;
        int right=height.size()-1;
        while(left<right){
            int width=right-left;
            int high=0;
            // 移动指针的逻辑，宽度更小了，主要是看高度，h=min(short,long) 若短边移动，可能更大；若长边移动，则不变或者更小，故最终面积不变或者更小故应该移动短边
            if(height[left]<height[right]){
                // 面积取更小的.
                high=height[left];
                left++;
            }else{
                high=height[right];
                right--;
            }
            if(res<width*high){
                res=width*high;
            }
        }
        return res;
    }
};
