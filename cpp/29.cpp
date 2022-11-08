class Solution {
    /*
    题目：给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符和long。
    以 99/5 为例子，对5不断翻倍，5 -> 10 -> 20 -> 40 -> 80 -> 160，此时160超出99，即可用99-80，剩余19留给下一轮。5 -> 80总共翻倍4次，也就是2^4=16倍，计数器+16，本轮结束； 　　重复以上操作，19/5，同理5 -> 10 -> 20，此时20超出19，即可用19-10，剩余9留给下一轮。5 -> 10总共翻倍1次，也就是2^1=2倍，计数器+2，本轮结束； 　　再次重复以上操作，9/5，同理5 -> 10，此时10超出9，即可用9-5，剩余4留给下一轮。5 -> 5总共翻倍0次，也就是2^0=1倍，计数器+1，本轮结束； 　　再次重复以上操作，4已经小于5，计算即可结束。 　　计数器16+2+1=19，即可得知99/5=19（省略小数部分）
    这里为了避免左移翻倍超时 直接改成右移
    */
public:
    int divide(int dividend, int divisor) {
        int res=0;
        int sigh=-1;
        if(divisor==int_min){
            // 除数边界值特殊处理
            if(dividend==int_min){
                return 1;
            }else {
                // 除数绝对着最大
                return 0;
            }
        }
        if(dividend==int_min){
            // 被除数边界值特殊处理
            if(divisor==-1){
                //overflow
                return int_max;
            }else if(divisor==1){
                return int_min;
            }
            dividend+=abs(divisor);
            res++;// 先执行一次加操作，避免abs转换溢出
        }
        if((dividend>0&&divisor>0)||(dividend<0&&divisor<0)){
            sigh=1;
        }
        int a=abs(dividend);
        int b=abs(divisor);
        while (a >= b) {
            int c = 1;
            int s = b;
            // 需指数级快速逼近，以避免执行超时
            while (s < (a >> 1)) { // 逼近至一半，同时避免溢出
                s += s; // *2
                c += c;// *2
            }
            res += c;
            a -= s;
        }
        // 因为不能使用乘号，所以将乘号换成三目运算符
        return sigh == 1 ? res : -res;
    }
    const int int_min=0x80000000;
    const int int_max=0x7FFFFFFF;
};
