class Solution {
    /*
    给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。
    [5,7]
    101
    110
    111
    可以看到只要有一位为0，最终该位置就为0，5-7范围内最终左起第一位全是1
    我们可以一起右移动位数，直到left>=right说明已经到了全是1的位
    故结果就是都是1的位+多个0
    */
public:
    int rangeBitwiseAnd(int left, int right) {
        int res=0;
        int zeros=0;
        while(left<right){
            left>>=1;
            right>>=1;
            zeros++;
        }
        return left<<zeros;
    }
};
