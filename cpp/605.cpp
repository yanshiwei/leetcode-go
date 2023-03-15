class Solution {
    /*
    假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给你一个整数数组  flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。
    */
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        // 贪心：从左到右遍历：能种花条件是 当前位置上0且前一个位置上0or边界并且后一个位置上0or边界
        // 遍历结束就是最多可以种花的结果
        int length=flowerbed.size();
        for(int i=0;i<length;i++){
            // 能种花条件是 当前位置上0且前一个位置上0or边界并且后一个位置上0or边界
            if(flowerbed[i]==0){
                if((i==0||flowerbed[i-1]==0)&&(i==length-1||flowerbed[i+1]==0)){
                    n--;
                    flowerbed[i]=1;
                }
            }
        }
        return n<=0;
    }
};
