class Solution {
/*
给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。
*/
public:
    int minDays(vector<int>& bloomDay, int m, int k) {
        // 每束花需要 k 朵花，需要制作 m 束花，因此一共需要 k×m朵花。如果花园中的花朵数量少于 k×mk ，即数组 bloomDay的长度小于 k×m，则无法制作出指定数量的花束，返回 −1。如果数组 bloomDay的长度大于或等于 k×m，则一定可以制作出指定数量的花束
        if(m>bloomDay.size()/k){
            //m*k will overflow
            return -1;
        }
        // 显而易见的结论是当 「数组里面花朵全部盛开」 的时候我们肯定可以制作所需的全部花束。数组的最大值这里我们不妨计作为 iii 。由于题目问的是符合要求的「最少天数」是多少。最大天数的最小化 二分
        // 1 范围 min max
        int left=1e9;
        for(auto &b:bloomDay){
            if(left>b){
                left=b;
            }
        }
        int right=0;
        for(auto &b:bloomDay){
            if(right<b){
                right=b;
            }
        }
        while(left<right){
            int mid=(right-left)/2+left;
            if(check(bloomDay,mid,m,k)){
                //满足制作m多花 在mid天 为了最小化
                right=mid;
            }else{
                left=mid+1;
            }
        }
        return left;
    }
private:
// 等待天数为mid时 能否制作m*k
    bool check(vector<int>& bloomDay,int mid, int m, int k){
        int numOfEach=0;//花朵数
        int numOfFlo=0;// 花束数
        for(auto&b:bloomDay){
            if(b<=mid){
                // 已经开放可以做花
                numOfEach++;
                if(numOfEach==k){
                    //说明可以制作一束了
                    numOfFlo++;
                    numOfEach=0;// 为下一束准备
                }
            }else{
                // 没有开放还不能做
                numOfEach=0;//不连续，直接清零
            }
        }
        return numOfFlo>=m;
    }
};
