/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
// //递归
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//     //肯定都在左子树
// 	if p.Val<root.Val&&q.Val<root.Val{
//         return lowestCommonAncestor(root.Left,p,q)
//     }
//     //肯定都在右子树
// 	if p.Val>root.Val&&q.Val>root.Val{
//         return lowestCommonAncestor(root.Right,p,q)
//     }
//     //说明这两个节点在以本次root节点为根节点
//     //的左右子树上,那么只有该节点可能是最近的公共祖先 
//     return root  
// }
//非递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
        //都在左子树
        if p.Val<root.Val&&q.Val<root.Val{
            root=root.Left
        }else if p.Val>root.Val&&q.Val>root.Val{
            root=root.Right
        }else{
            break
        }
    }
    return root
}
