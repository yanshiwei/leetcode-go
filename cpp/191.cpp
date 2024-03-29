class Solution {
    // 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
public:
    int hammingWeight(uint32_t n) {
        int res=0;
        while(n){
            if(n&1==1) res++;
            n>>=1;//n每次向右移动一位，即除以2
        }
        return res;
    }
};
