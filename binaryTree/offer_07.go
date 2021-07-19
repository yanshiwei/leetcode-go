/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return nil
    }
    // create root
    root:=&TreeNode{Val:preorder[0]}
    // 获取根节点在中序遍历数组中的index
    var rootIdx int
    for rootIdx<len(inorder){
        if inorder[rootIdx]==root.Val{
            break
        }
        rootIdx++
    }
    // 递归，左子树长度rootIdx
    root.Left=buildTree(preorder[1:rootIdx+1],inorder[0:rootIdx])
    root.Right=buildTree(preorder[rootIdx+1:],inorder[rootIdx+1:])
    return root
}
