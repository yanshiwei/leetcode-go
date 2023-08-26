class Solution {
    /*
    给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
    */
public:
    int findMaxForm(vector<string>& strs, int m, int n) {
        // 题目包括2个背包，一个是1的数目总数不超过m，另一个是0的数目不超过n
        // 题目转化为01背包问题：范围strs，重量是strs[k]中1的数目/0的数目，价值都是1（选这个字符串长度+1）
        // dp[k][i][j] 代表考虑前 k 件物品，在数字 0 容量不超过 i，数字 1 容量不超过 j的条件下的「最大价值」（每个字符串的价值均为 1）
        // 故dp[k][i][j]=max(dp[k-1][i][j],dp[k][i-cnt[k][0]][j-cnt[k][1]]+1)
        // 其中cnt[k] 数组记录的是字符串k中出现的 01 数量
        // 然后降维度，dp[i][j]=max(dp[[i[j],dp[i-cnt[k][0]][j-cnt[k][1]]+1)
        int length=strs.size();
        // 求每个字符包含的0 1 数量
        vector<vector<int>>cnt(length,vector<int>(2,0));
        for(int i=0;i<strs.size();i++){
            string str=strs[i];
            int zero=0;
            int one=0;
            for(char c: str){
                if(c=='0'){
                    zero++;
                }else{
                    one++;
                }
            }
            cnt[i][0]=zero;
            cnt[i][1]=one;
        }
        vector<vector<int>>dp(m+1,vector<int>(n+1,0));//dp[i][j]在数字 0 容量不超过 i，数字 1 容量不超过 j的条件下的最大长度
        //// 原始dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
        // dp[k][i][j]=max(dp[k-1][i][j],dp[k][i-cnt[k][0]][j-cnt[k][1]]+v[k])
        // 这里2个01背包最值问题，v[k]=1,w[k]=cnt[k][0] or cnt[k][1]
        for(int k=1;k<=length;k++){
            int zero=cnt[k-1][0];//第k个字符串包含的0个数
            int one=cnt[k-1][1];//第k个字符串包含的1个数
            // 01背包 逆序
            for(int i=m;i>=zero;i--){
                for(int j=n;j>=one;j--){
                    dp[i][j]=max(dp[i][j],dp[i-zero][j-one]+1);
                }
            }
        }
        return dp[m][n];
    }
};
