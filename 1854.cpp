class Solution {
    /*
    给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。
年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1] 内。注意，人不应当计入他们死亡当年的人口中。
返回 人口最多 且 最早 的年份。
1 <= logs.length <= 100
1950 <= birthi < deathi <= 2050
由于数据范围很小，所以可以直接暴力。从年份开始循环，然后再循环logs数组，如果年份在出生和死亡年份之间就记录。
    */
public:
    int maximumPopulation(vector<vector<int>>& logs) {
        int res=0;
        int cnt=0;
        for(int i=1950;i<=2050;i++){
            int cur=0;//记录当前年份存活人数
            for(auto log:logs){
                if(i>=log[0]&&i<log[1]){
                    cur++;//[birthi, deathi - 1] 
                }
            }
            if(cur>cnt){
                //记录存活人数最大值
                res=i;
                cnt=cur;
            }
        }
        return res;
    }
};
