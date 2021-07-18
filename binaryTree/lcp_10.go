/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimalExecTime(root *TreeNode) float64 {
    /*
https://leetcode-cn.com/problems/er-cha-shu-ren-wu-diao-du/solution/dfs-si-lu-dai-ma-he-zheng-ming-by-leetcode-solutio/
    双核CPU可以考虑左右两边在CPU时间上使用双核（注意是cpu时间而不是同一个任务在两个核上执行）
    设 f(i) 为节点 i 的最短时间
    1.两个子树分别使用双核跑完，然后再根节点跑。这样的时间消耗是 f(l)+f(r)+val(i)
    2.在左边花费 x 时间使用双核（这里是有效使用，即两个核都必须用上，后同），在右边花费 yy 时间使用双核，然后在左右两棵子树一边一个核。
    这种情况下，设左右子树的任务总时间和分别为 sum(l),sum(r)。则通过利用双核，左子树剩下的任务时间变成了 sum(l) - 2x，右子树剩下的任务时间变成了 sum(r) - 2y，然后一边一个核处理完剩下的任务所需时间为这两者的较大者股总时间
    val(i)+x+y+max{sum(l)−2x,sum(r)−2y}
    */
    tc,pc:=dfs(root)
    return tc-pc
}

func dfs(root *TreeNode)(float64,float64){
    if root==nil{
        return 0,0
    }
    leftSum,maxConLeft:=dfs(root.Left)
    rightSum,maxConright:=dfs(root.Right)
    tc:=float64(root.Val)+leftSum+rightSum//total time
    //只考虑 leftSum >= rightSum 的情况，不满足就交换
    if leftSum<rightSum{
        leftSum,rightSum=rightSum,leftSum
        maxConLeft,maxConright=maxConright,maxConLeft
    }
    var pc float64//parallel time
    if leftSum-2*maxConLeft<=rightSum{
        pc=(leftSum+rightSum)/2
    }else{
        pc=rightSum+maxConLeft
    }
    return tc,pc
}

func min(x, y float64)float64{
    if x<y {
        return x
    }
    return y
}


func max(x, y float64)float64{
    if x<y {
        return y
    }
    return x
}
