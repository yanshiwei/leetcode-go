/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var leftNum int
 var rightNum int
 var value int
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    /*
    1.一个节点，如果取它的左（右）节点，则左（右）节点的所有节点都是你的。取父节点，则父节点以上的所有节点也都是你的
    2.故取胜的关键是取得一半以上的节点数量，另外数字是无序的，随便填写，所以不考虑数字顺序。总共有n个节点（n为奇数），考虑其一半m = n / 2
    */
    leftNum=0
    rightNum=0
    value=x
    dfs(root)
    var half=n/2
    if leftNum>half||rightNum>half{
        //选择左或右子树，节点数均大于half
        return true
    }
    if leftNum+rightNum<half{
        //选择父节点，节点数大于half
        return true
    }
    //left==half || right==half 无论选取哪个节点（左/右/父），其他另一边节点和父节点都是一号玩家的，一号玩家总数超过一半，都是必输
    //half<=left+right<n,无论选取哪个节点（左/右/父）,其他另一边节点和父节点都是一号玩家的，一号玩家总数超过一半，都是必输
    return false
}
//该节点能提供的节点个数（左右子树+本身）
func dfs(root *TreeNode)int{
    if root==nil {
        return 0
    }
    left:=0
    right:=0
    left=dfs(root.Left)
    right=dfs(root.Right)
    if root.Val==value{
        leftNum=left
        rightNum=right
    }
    return left+right+1
}
