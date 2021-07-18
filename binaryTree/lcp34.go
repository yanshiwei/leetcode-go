/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxValue(root *TreeNode, k int) int {
    /*
    每个节点可以有两种处理结果：染色或者不染色;
    根据题意显然是要用后序遍历，自底向上，由下层给出结果，供上层使用和选择
    定义：
    dp[0] 表示：以某个节点为根节点，根节点不染色的情况:
    dp[0] = 左子节点处理的最大值+右子节点处理的最大值
    dp[i] 表示：以某个节点为根节点，（根节点染色的情况），蓝色相连节点为i个的最大染色值， 1 <= i <= k，故dp长度为k+1
    dp[i] = leftDp[j] + rightDp[i-1-j] + root.Val     (i = [1,k] , j=[0,i）)
    */
    dp:=dfs(root,k)
    return getMax(dp)//统一比较染色和不染色情况
}

func dfs(root *TreeNode, k int)[]int {
    var dp =make([]int,k+1)
    if root==nil{
        return dp
    }
    leftDp:=dfs(root.Left,k)
    rightDp:=dfs(root.Right,k)
    //不染色情况
    dp[0]=getMax(leftDp)+getMax(rightDp)
    //染色情况
    for i:=1;i<=k;i++{
        //一共i个
        for j:=0;j<i;j++{
            //左子树分配j个
            dp[i]=max(dp[i],leftDp[j]+rightDp[i-1-j]+root.Val)
        }
    }
    return dp
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

func getMax(dp []int)int {
    var res=dp[0]
    for i:=1;i<len(dp);i++{
        if dp[i]>res{
            res=dp[i]
        }
    }
    return res
}
