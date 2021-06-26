/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    var inorder []int
    for i:=range preorder{
        inorder=append(inorder,preorder[i])
    }
    sort.Ints(inorder)
    return buildTreeFromPreAndIn(preorder,inorder)
}

func buildTreeFromPreAndIn(preorder[]int,inorder[]int)*TreeNode{
    if len(preorder)==0{
        return nil
    }
    var rootIdx int
    // 1. 创建根节点
    root:=&TreeNode{Val:preorder[0]}
    // 2. 获取根节点在中序遍历数组中的index
    for rootIdx<len(inorder){
        if inorder[rootIdx]==root.Val{
            break
        }
        rootIdx++
    }
    // 3. 递归。左子树长rootIdx
    root.Left=buildTreeFromPreAndIn(preorder[1:rootIdx+1],inorder[:rootIdx])
    root.Right=buildTreeFromPreAndIn(preorder[rootIdx+1:],inorder[rootIdx+1:])
    return root
}
