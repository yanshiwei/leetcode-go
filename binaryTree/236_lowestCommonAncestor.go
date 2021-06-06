/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     /*
     过递归对二叉树进行后序遍历，当遇到节点 p 或 q时返回。从底至顶回溯，当节点p,q 在节点 root的异侧时，节点 root即为最近公共祖先，则向上返回 root
     参考：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
     */
     if root==nil{
         return nil
     }
     //当 root 等于 p, q，则直接返回 root
     if root==p||root==q{
         return root
     }
     left:=lowestCommonAncestor(root.Left,p,q)
     right:=lowestCommonAncestor(root.Right,p,q)
     //同时为空 ：说明 root的左 / 右子树中都不包含p,q ，返回 null
     if left==nil&&right==nil{
         return nil
     }
     //只有right不为空，说明p,q不在root的左侧，直接返回right
     if left==nil&&right!=nil{
         return right
     }
     //只有left不为空，说明p,q不在root的右侧，直接返回left
     if left!=nil&&right==nil{
         return left
     }
    //同时不为空 ：说明 p,q 分列在 root 的 异侧 （分别在 左 / 右子树），因此 root 为最近公共祖先，返回 root
     return root
}
