class Solution {
public:
    int maxDistance(vector<int>& position, int m) {
        // 最小值最大化 二分
        // 1 确定范围 任意两个球之间最小磁力，最小是0，最大
        int left=1,right=1e9;
        sort(position.begin(),position.end());
        right=position.back()-position.front();//vector: front back push_back pop_back
        while(left<right){
            int mid = ((right - left)/2) + left;/* 
            不使用：
            int mid=left+(right-left)/2;
            是因为left 可能不变。比如左右边界为0 1，此时 mid = (0 + 1) / 2 -> 0，故产生死循环
            故mid = left + (right - left + 1) / 2;
            */
            if(check(position,mid,m)){
                //任意两个球之间最小磁力mid 能放得下m个球 为了最大化
                left=mid;
            }else{
                //任意两个球之间最小磁力mid 不能放得下m个球
                right=mid-1;
            }
        }
        return left;
    }
private:
    // m个球 任意两个球之间最小磁力mid时，也即小球的间距大于等于 mid满足这个条件后的球个数能否达到m
    bool check(vector<int>& position, int mid, int m){
        // 要求任意两个球之间磁力>=mid，则我们预先对给定的篮子的位置进行排序，那么从贪心的角度考虑，第一个小球放置的篮子一定是 position 最小的篮子，即排序后的第一个篮子。那么为了满足上述条件，第二个小球放置的位置一定要大于等于 position[0]+mid，接下来同理。从前往后扫 position 数组，看在当前答案mi下我们最多能在篮子里放多少个小球，我们记这个数量为 cnt，如果 cnt 大于等于 m，那么说明当前答案下我们的贪心策略能放下 m 个小球且它们间距均大于等于mid
        int pre=position[0];
        int cnt=1;// 组数（球数）
        for(int i=1;i<position.size();i++){
            if(position[i]-pre>=mid){
                // 排序后不改变磁力计算的结果 也方便我们求磁力 直接减去即可
                pre=position[i];
                cnt++;
            }
        }
        return cnt>=m;
    }
};
