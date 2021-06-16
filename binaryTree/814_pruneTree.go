/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    //从底到上，故后序遍历
    if dfs(root){
        return root
    }else{
        //不包含1则置空
        return nil
    }
}
//判断以 node 为根的子树中是否包含 1，其不包含 1 当且仅当以 node 的左右孩子为根的子树均不包含 1，并且 node 节点本身的值也不为 1
func dfs(node *TreeNode)bool{
    if node==nil{
        //不包含1
        return false
    }
    left:=dfs(node.Left)
    right:=dfs(node.Right)
    //如果 node 的左右孩子为根的子树不包含 1，那我们就需要把对应的指针置为空
    if left==false{
        node.Left=nil
    }
    if right==false{
        node.Right=nil
    }
    return node.Val==1|| left||right
}
