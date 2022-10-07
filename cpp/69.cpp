class Solution {
    /*
    给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
    */
public:
    int mySqrt(int x) {
        if(x<=1){
            return x;
        }
        int i=0;
        while(i<=x/2&&(long)i*i<=x){
            i++;
        }
        return i-1;
    }
};
