/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
 定义一个递归方法，从下到上递归，后续遍历。
 1. 每个节点的贡献：
    1.1 空节点的最大贡献值等于 0；
    1.2 非空节点的最大贡献值等于节点值与其子节点中的最大贡献值之和
 2. 每个接的的最大路径和：
    该节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值，如果子节点的最大贡献值为正，则计入该节点的最大路径和，否则不计入该节点的最大路径和。维护一个全局变量 maxSum 存储最大路径和，在递归过程中更新 maxSum 的值，最后得到的 maxSum 的值即为二叉树中的最大路径和
 参考：
    1.https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/solution/cchao-jian-dan-li-jie-miao-dong-by-fengz-8l3m/
    2.https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/solution/er-cha-shu-zhong-de-zui-da-lu-jing-he-by-leetcode-/ 
 */
var maxV int
func maxPathSum(root *TreeNode) int {
    if root==nil{
        return 0
    }
    maxV=INTMIN
    dfs(root)
    return maxV
}
//通过后序遍历的方式，先计算出左右子树的最大路径和，然后再计算当前树的最大路径和
func dfs(root *TreeNode)int {
    if root==nil{
        return 0
    }
    //计算左边分支最大值，左边分支如果为负数还不如不选择
    leftV:=max(0,dfs(root.Left))
    //计算右边分支最大值，右边分支如果为负数还不如不选择
    rightV:=max(0,dfs(root.Right))
    //一个子树内部的最大路径和 = 左子树提供的最大路径和 + 根节点值 + 右子树提供的最大路径和，作为路径与已经计算过历史最大值做比较
    maxV=max(maxV,root.Val+leftV+rightV)
    //返回当前子树对外提供的贡献
    return root.Val+max(leftV,rightV)
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
