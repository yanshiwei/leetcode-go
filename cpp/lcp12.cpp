class Solution {
    /*
    为了提高自己的代码能力，小张制定了 LeetCode 刷题计划，他选中了 LeetCode 题库中的 n 道题，编号从 0 到 n-1，并计划在 m 天内按照题目编号顺序刷完所有的题目（注意，小张不能用多天完成同一题）。
在小张刷题计划中，小张需要用 time[i] 的时间完成编号 i 的题目。此外，小张还可以使用场外求助功能，通过询问他的好朋友小杨题目的解法，可以省去该题的做题时间。为了防止“小张刷题计划”变成“小杨刷题计划”，小张每天最多使用一次求助。
我们定义 m 天中做题时间最多的一天耗时为 T（小杨完成的题目不计入做题总时间）。请你帮小张求出最小的 T是多少。
示例 1：
输入：time = [1,2,3,3], m = 2
输出：3
解释：第一天小张完成前三题，其中第三题找小杨帮忙；第二天完成第四题，并且找小杨帮忙。这样做题时间最多的一天花费了 3 的时间，并且这个值是最小的。
    */
public:
    int minTime(vector<int>& time, int m) {
        // 最大值的最小化，二分
        // 1 确定范围，m 天中做题时间最多的一天耗时，最小为0，最大sum(time)
        int left=0;
        int right=0;
        for(auto&d:time){
            right+=d;
        }
        while(left<right){
            int mid=left+(right-left)/2;
            if(check(time,mid,m)){
                // m 天中做题时间最多的一天耗时mid时可以满足 为了最小化mid
                right=mid;// mid满足故包括mid
            }else{
                left=mid+1;
            }
        }
        return left;
    }
private:
    // 验证m 天中做题时间最多的一天耗时为mid时能不能m天做完，也即每组累加和在<=mid的情况下 能否在m组内分完
    bool check(vector<int>& time, int mid, int m){
        int cnt=1;//cnt为分的总组数
        int rest=mid;//rest为当前组组剩余容量
        int maxx=-1;// 当天当前作业最大耗时
        bool flag=true;// 可以求助
        for(int i=0;i<time.size();i++){
            maxx=max(maxx,time[i]);//维护当前组（也就是当天）的最大值
            if(rest>=time[i]){
                //能装下时就直接装
                rest-=time[i];
            }else if(flag){
                //装不下且可以求助时，把当前的最费时的那个拿去求助
                flag=false;
                rest+=maxx;
                i--;// 继续当前作业，下一次就可以直接装
            }else{
                //装不下 且无法求助时 只能从第二天开始了(开始下一组)
                cnt++;// 新的一组
                maxx=-1;// 更新当天最大值
                flag=true;// 更新当前求助
                rest=mid;
                i--;// 继续当前作业
            }
        }
        return cnt<=m;//m组内能分完即可(=m天内能看完)
    }
};
