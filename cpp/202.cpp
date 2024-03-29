class Solution {
    /*
    快乐数」 定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
输入：n = 19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
    如果没有循环，则最终得到1
    如果循环，则得不到1，怎么判断有无循环，只需要记录之前计算的值有没有再出现即可
    */
public:
    bool isHappy(int n) {
        unordered_set<int>hash;
        while(1){
            int next=getSquareNum(n);
            if(next==1){
                return true;
            }
            if(hash.count(next)>0){
                break;
            }
            hash.insert(next);
            n=next;
        }
        return false;
    }
private:
    int getSquareNum(int n){
        int res=0;
        while(n>0){
            int t=n%10;
            res+=t*t;
            n/=10;
        }
        return res;
    }
};
