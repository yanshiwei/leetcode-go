class Solution {
    // 颠倒给定的 32 位无符号整数的二进制位。
public:
    uint32_t reverseBits(uint32_t n) {
        int res=0;
        for(int i=0;i<32;i++){
            res<<=1;//左移动空出0
            res|=(n&1);
            n>>=1;//右移动
        }
        return res;
    }
};
