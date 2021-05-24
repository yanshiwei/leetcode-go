func numTrees(n int) int {
    if n==1{
        return 1
    }
    /*
    动态规划
    1.假设 n 个节点存在二叉排序树的个数是 G (n)
    令 f(i) 为以 i 为根的二叉搜索树的个数，则
    G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)
    2.当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i则
    f(i)=G(i−1)∗G(n−i)
    3. 综合两个可得：
    G(n)=G(0)∗G(n−1)+G(1)∗G(n−2)+...+G(n−1)∗G(0)
    故G(2)=G(0)*G(1)+G(1)*(0)
    G(3)=G(0)*G(2)+G(1)*G(1)+G(2)*G(0)
    ...
    G(i)=G(0)*G(i-1)+G(1)*G(i-2)+G(2)*G(i-3)+...+G(i-1)*(0)
    */
    var dp=make([]int,n+1)
    dp[0]=1//初始化
    dp[1]=1//init
    for i:=2;i<n+1;i++{
        //i个节点时
        for j:=1;j<i+1;j++{
            dp[i]+=dp[j-1]*dp[i-j]
        }
    }
    return dp[n]
}
