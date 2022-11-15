class Solution {
public:
    int trap(vector<int>& height) {
        //给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
        if(height.size()<1){
            return 0;
        }
        int left=0,right=height.size()-1;
        int left_max=0,right_max=0;
        int res=0;
        while(left<=right){
            if(height[left]<height[right]){
                // left_max height[left]更新 故left_max<right_max 右侧构成一个墙壁
                if(height[left]<left_max){
                    // 可以蓄水
                    res+=left_max-height[left];
                }else{
                    left_max=height[left];
                }
                left++;
            }else{
                // right_max 由height[right]更新 故left_max>right_max 左侧构成一个墙壁
                if(height[right]<right_max){
                    res+=right_max-height[right];
                }else{
                    right_max=height[right];
                }
                right--;
            }
        }
        return res;
    }
};
