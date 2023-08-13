class Solution {
    /*
    给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量
    */
public:
    int maxArea(vector<int>& height){
        /*
        s=min(h[left],h[right])*(right-left)
        每走一步，都会导致容器宽度减1
        1.若向内移动短板，水槽的短板 min(h[i],h[j])可能变小也可能变大，因此下个水槽的面积 可能增大；
        2.若向内移动长板，水槽的短板 min(h[i],h[j])一定是不变或者变小，因为长板移动要么更大，此时短板还是原来值；要么不变，此时短板还是原来值，要么变小，比原来短板值更小此时就变小了。因此下个水槽的面积 一定变小
        因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积
        */
        int n=height.size();
        if(n<1){
            return 0;
        }
        int res=0;
        // 相向指针
        int left=0;
        int right=n-1;
        while(left<right){
            // 两者相遇不能构成容器，是结束条件，故不需要=
            int width=right-left;
            int high=0;
            // 移动短板才可能更大
            if(height[left]<height[right]){
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
/*
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
    */
};
