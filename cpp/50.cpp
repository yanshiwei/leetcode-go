class Solution {
public:
    double myPow(double x, int n) {
        if(x==0){
            return 0;
        }
        //int32变量n∈[−2147483648,2147483647] ，因此当 n = -2147483648时执行 n = -n会因越界而赋值出错。解决方法是先将 n 存入 long 变量 N后面用N 操作即可
        long N=n;
        return N >= 0 ? powCore(x, N) : 1.0 / powCore(x, -N);
    }
private:
    double powCore(double x, long n){
        if(n==0){
            return 1.0;
        }
        double y=powCore(x,n/2);
        if((n&1)==0){
            // 偶数
            return y*y;
        }else{
            // 奇数
            return y*y*x;
        }
    }
};
