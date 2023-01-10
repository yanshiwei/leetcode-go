class Solution {
public:
    /*
    令dp[i][j]为字符子串s1[0, i),s2[0, j)能否组成s3[0, i+j)。 
    ps: s1[0, i) = s1.substring(0, i),不包含i。
    对于dp[i][j]该状态来说，要想组成s3[0,i+j)，其s3[0, i+j)最后一个字符s3[i+j-1]要么来自s1[i-1], 要么来自s2[j-1],因此，状态转移：
    1. 若s1[i-1]==s3[i+j-1]:dp[i][j]=dp[i-1][j],i>0;
    2. 若s2[j-2]==s3[i+j-1]:dp[i][j]=dp[i][j-1],j>0;
    初始化dp[0][0]=true表示两个空字符串能够组成一个空字符串
    */
    bool dfs(string s1,int i,string s2,int j,string s3,int k){
        //长度不一致，直接返回
        if(s1.size()+s2.size()!=s3.size()){
            return false;
        }
        // i j k都正好遍历完
        if(i==s1.size()&&j==s2.size()&&k==s3.size()){
            return true;
        }
        if(k>=s3.size()){
            return false;
        }
        // i没有遍历完，如s1[i]==s3[k] 尝试迭代
        if(i<s1.size()){
            if(s1[i]==s3[k]){
                if(dfs(s1,i+1,s2,j,s3,k+1)){
                    return true;
                }
            }
        }
        // 如果i迭代失败，回溯到j
        if(j<s2.size()){
            if(s2[j]==s3[k]){
                if(dfs(s1,i,s2,j+1,s3,k+1)){
                    return true;
                }
            }
        }
        return false;               
    }
    // dfs原始方法超时，用空间换取时间。引入map;
    bool dfsBetter(string s1,int i,string s2,int j,string s3,int k,unordered_map<string,int>&hash){
        //长度不一致，直接返回
        if(s1.size()+s2.size()!=s3.size()){
            return false;
        }
        //key 的话用字符串 i + "@" + j ，之所以中间加 "@"，是为了防止 i = 1 和 j = 22。以及 i = 12，j = 2。这样的两种情况产生的就都是 122。加上 "@" 可以区分开来。
        string key=to_string(i)+"@"+to_string(j);
        // int to string to_string; string to int atoi(str.c_str());
        if(hash.count(key)){
            //之前已经遍历过i/j
            return hash[key]==1;
        }
        // i j k都正好遍历完
        if(i==s1.size()&&j==s2.size()&&k==s3.size()){
            hash[key]=1;
            return true;
        }
        if(k>=s3.size()){
            hash[key]=0;
            return false;
        }
        // i没有遍历完，如s1[i]==s3[k] 尝试迭代
        if(i<s1.size()){
            if(s1[i]==s3[k]){
                if(dfs(s1,i+1,s2,j,s3,k+1)){
                    return true;
                }
            }
        }
        // 如果i迭代失败，回溯到j
        if(j<s2.size()){
            if(s2[j]==s3[k]){
                if(dfs(s1,i,s2,j+1,s3,k+1)){
                    return true;
                }
            }
        } 
        hash[key]=0;
        return false;               
    }
    bool isInterleave(string s1, string s2, string s3) {
        // return dfs(s1,0,s2,0,s3,0);
        // unordered_map<string,int>hash;
        // return dfsBetter(s1,0,s2,0,s3,0,hash);
        int n=s1.size();
        int m=s2.size();
        if(n+m!=s3.size()){
            return false;
        }
        vector<vector<bool>>dp(n+1,vector<bool>(m+1,false));
        for(int i=0;i<=n;i++){
            for(int j=0;j<=m;j++){
                if(i==0&&j==0){
                    dp[i][j]=true;//init
                }else{
                    if(i>0&&s1[i-1]==s3[i+j-1]){
                        dp[i][j]=dp[i][j]||dp[i-1][j];
                    }
                    if(j>0&&s2[j-1]==s3[i+j-1]){
                        dp[i][j]=dp[i][j]||dp[i][j-1];
                    }                   
                }
            }
        }
        return dp[n][m];
    }
};
