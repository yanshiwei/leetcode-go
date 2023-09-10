class Solution {
public:
    int getSum(int n){
        int sum=0;
        while(n){
            sum+=(n%10)*(n%10);
            n/=10;
        }
        return sum;
    }
    bool isHappy(int n) {
        int sum=0;
        unordered_set<int>sset;
        while(1){
            sum=getSum(n);
            if(sum==1){
                return true;
            }
            if(sset.count(sum)){
                // 如果这个sum曾经出现过，说明已经陷入了无限循环了，立刻return false
                return false;
            }
            sset.insert(sum);
            n=sum;
        }
        return true;
    }
};
