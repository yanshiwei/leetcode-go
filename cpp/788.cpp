class Solution {
public:
/*
我们称一个数 X 为好数, 如果它的每位数字逐个地被旋转 180 度后，我们仍可以得到一个有效的，且和 X 不同的数。要求每位数字都要被旋转。
如果一个数的每位数字被旋转以后仍然还是一个数字， 则这个数是有效的。0, 1, 和 8 被旋转后仍然是它们自己；2 和 5 可以互相旋转成对方（在这种情况下，它们以不同的方向旋转，换句话说，2 和 5 互为镜像）；6 和 9 同理，除了这些以外其他的数字旋转以后都不再是有效的数字
    一个数是好数：
    1 数中必须没有出现3 4 7 is -1
    2 数中至少出现一次2 或 5 或 6 或9 is 1
    3 对于0 1 8没有要求 is 0
*/
    int rotatedDigits(int n) {
        int cnt=0;
        for(int i=1;i<=n;i++){
            string num_str=to_string(i);
            bool valid=true;
            bool hasDiff=false;
            for(char ch:num_str){
                if(check[ch-'0']==-1){
                    // 含有3 4 7
                    valid=false;
                }else if(check[ch-'0']==1){
                    // 至少出现一次2 或 5 或 6 或9
                    hasDiff=true;
                }
            }
            if(valid&&hasDiff){
                cnt++;
            }
        }
        return cnt;
    }
    int check[10]={0,0,1,-1,-1,1,1,-1,0,1};
};
