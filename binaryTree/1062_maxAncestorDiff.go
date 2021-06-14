/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func maxAncestorDiff(root *TreeNode) int {
    res=0
    dfs(root,root.Val,root.Val)
    return res
}
//就一个节点来说所谓最大差值，就是祖先的最大值或者最小值和自己的val的差值。
//只需要知道所有祖先可能的最大值和最小值，在遍历时携带传递即可。
//pre order
func dfs(root *TreeNode,big,small int){
    if root==nil{
        return
    }
    res=max(res,max(abs(root.Val-big),abs(root.Val-small)))
    //递归更新祖先的最值
    big=max(big,root.Val)
    small=min(small,root.Val)
    dfs(root.Left,big,small)
    dfs(root.Right,big,small)
}

func abs(x int)int {
    if x<0 {
        return -1*x
    }
    return x
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
