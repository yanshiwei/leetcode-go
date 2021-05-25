/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    /*
    如果该二叉树的左子树不为空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值；它的左右子树也为二叉搜索树。
    */
    return helper(root,INTMIN,INTMAX)
}
//函数表示考虑以 root 为根的子树，判断子树中所有节点的值是否都在（min,max） 的范围内。如果 root 节点的值 val 不在 (min,max) 的范围内说明不满足条件直接返回，否则我们要继续递归调用检查它的左右子树是否满足，如果都满足才说明这是一棵二叉搜索树

func helper(root *TreeNode,min,max int)bool{
    if root==nil{
        return true
    }
    if root.Val<=min || root.Val>=max{
        //root 节点的值 val 不在 (min,max) 的范围内说明不满足条件直接返回
        return false
    }
    return helper(root.Left,min,root.Val)&&helper(root.Right,root.Val,max)
}
const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
