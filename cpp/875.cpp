class Solution {
    /*
    珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）
    */
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        // （要吃掉所有香蕉速度比较较大）最大值的最小化 二分
        // 1. 范围
        int left=1;// 速度最小的时候，耗时最长
        int right=0;// max of arr
        for(int pile:piles){
            if(right<pile){
                right=pile;
            }
        }
        while(left<right){
            int mid=left+(right-left)/2;//吃香蕉最大速度
            if(check(piles,mid,h)==false){
                // 耗时太多，说明速度太慢了，要加大速度
                left=mid+1;
            }else{
                // h小时内mid速度可以吃完，为了保证最小
                right=mid;
            }
        }
        return left;
    }
private:
// 最大速度mid时能否h小时内全吃完
//珂珂一个小时之内只能选择一堆香蕉吃，因此：每堆香蕉吃完的耗时 = 这堆香蕉的数量 / 珂珂一小时吃香蕉的数量(向上取整)
//小技巧: 向下取整：a//b, 向上取整: (a + b - 1) // b
    bool check(vector<int>& piles,int mid, int h){
        int res=0;
        for(auto pile:piles){
            res+=(pile+mid-1)/mid;
        }
        return res<=h;
    }
};
