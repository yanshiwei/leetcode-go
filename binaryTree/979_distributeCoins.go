/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
func distributeCoins(root *TreeNode) int {
    res=0
    dfs(root)
    return res
}
//某个节点过载量=abs(节点上硬币数目-1),过载量就是需要将这个节点中的 金币移动到别的地方去或从其它地方移动进来
//dfs为寇哥节点所在子树的过载量=子树中金币的数量减去这个子树中节点的数量
func dfs(root *TreeNode)int{
    if root==nil{
        return 0
    }
    left:=dfs(root.Left)
    right:=dfs(root.Right)
    //fmt.Println(root.Val,left,right)
    res+=abs(left)+abs(right)//需要移动的次数
    return left+right+(root.Val-1)
}

func abs(x int)int{
    if x<0{
        return -1*x
    }
    return x
}
