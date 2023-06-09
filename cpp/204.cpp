class Solution {
    // 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
    // 从质数2开始，倍数增加4，6，8，都不是质数
    //下一个质数3，倍数增加9，12,注意每次找当前素数 x 的倍数时，是从x^2
    // 开始的。因为如果 x>2那么 2∗x肯定被素数 2 给过滤了
    // 以此类推
public:
    int countPrimes(int n) {
        if(n<3){
            return 0;
        }
        vector<bool> d(n, true);
        int res=0;
        for(long i=2;i<n;i++){
            if(d[i]){
                res++;
                // 该质数倍数增加,x^2开始
                for(long j=i*i;j<n;j+=i){
                    d[j]=false;
                }
            }
        }
        return res;
    }
};
