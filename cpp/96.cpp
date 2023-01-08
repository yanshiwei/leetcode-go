class Solution {
    /*
    给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
    dfs:time limit
    dp:
    1 假设n个节点存在的BST的个数是g(n)
    令f(i)为以i为根的BST的个数
    则
    g(n)=f(1)+f(2)+f(3)+...+f(n)
    2.当i为根结点时，左子树节点个数i-1个，右子树节点个数n-i个则
    f(i)=g(i-1)*g(n-i)
    3.综合两个可得：
    g(n)=g(0)*g(n-1)+g(1)*g(n-2)+...+g(n-1)*g(0)
    故有
    g(2)=g(0)*g(1)+g(1)*g(0)
    g(3)=g(0)*g(2)+g(1)*g(1)+g(2)*g(0)
    ...
    g(i)=g(0)*g(i-1)+g(1)+g(i-2)+...+g(i-1)*g(0)
    */
public:
    int dpfun(int n){
        if(n<1){
            return 0;
        }
        if(n==1){
            return 1;
        }
        vector<int>dp(n+1,0);//dp[i]代表i个节点存在的BST的个数
        dp[0]=1;
        dp[1]=1;
        for(int i=2;i<=n;i++){
            //i个节点时
            for(int j=1;j<=i;j++){
                //以j为根节点，j取值[1,i]
                dp[i]+=dp[j-1]*dp[i-j];
            }
        }
        return dp[n];
    }
    int dfs(int start,int end){
        if(start>end){
            return 1;
        }
        if(start==end){
            return 1;
        }
        int res=0;
        for(int i=start;i<=end;i++){
            auto left=dfs(start,i-1);
            auto right=dfs(i+1,end);
            res+=left*right;
        }
        return res;
    }
    int numTrees(int n) {
        return dpfun(n);
        // if(n<1){
        //     return 0;
        // }
        // return dfs(1,n);//time limit
    }
};
